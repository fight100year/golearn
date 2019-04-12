package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	c := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		defer close(c)
		w.Wait()
	}()

	select {
	case <-c:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("需要一个参数")
		return
	}

	var w sync.WaitGroup
	w.Add(1)

	t, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	duration := time.Duration(int32(t)) * time.Millisecond
	if timeout(&w, duration) {
		fmt.Println("time out")
	} else {
		fmt.Println("ok")
	}

	w.Done()
	if timeout(&w, duration) {
		fmt.Println("time out")
	} else {
		fmt.Println("ok")
	}
}
