package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	url   string
	delay = 5
	w     sync.WaitGroup
)

type data struct {
	r   *http.Response
	err error
}

func connect(c context.Context) error {
	defer w.Done()
	msg := make(chan data, 1)

	tran := &http.Transport{}
	client := &http.Client{Transport: tran}
	req, _ := http.NewRequest("GET", url, nil)

	go func() {
		response, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			msg <- data{nil, err}
			return
		}

		msg <- data{response, err}
	}()

	select {
	case <-c.Done():
		tran.CancelRequest(req)
		<-msg
		fmt.Println("request is canceled")
		return c.Err()
	case ok := <-msg:
		err := ok.err
		resp := ok.r

		if err != nil {
			fmt.Println("error select:", err)
			return err
		}
		defer resp.Body.Close()

		httpData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error select:", err)
			return err
		}
		fmt.Printf("server response: %s\n", httpData)
	}

	return nil
}

func main() {
	if len(os.Args) == 1 {
		return
	}

	url = os.Args[1]
	if len(os.Args) >= 3 {
		d, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		delay = d
	}

	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
	defer cancel()

	w.Add(1)
	go connect(c)
	w.Wait()
	fmt.Println("done")
}
