package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type client struct {
	id int
	i  int
}

type data struct {
	job    client
	square int
}

var (
	size    = 10
	clients = make(chan client, size)
	datas   = make(chan data, size)
)

func worker(w *sync.WaitGroup) {
	for c := range clients {
		datas <- data{c, c.i * c.i}
		time.Sleep(time.Second)
	}
	w.Done()
}

func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(datas)
}

func create(n int) {
	for i := 0; i < n; i++ {
		clients <- client{i, i}
	}
	close(clients)
}

func main() {
	fmt.Println(cap(clients), cap(datas))

	if len(os.Args) != 3 {
		fmt.Println("usage: xx #jobs #workers")
		return
	}

	jobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	workers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	go create(jobs)
	done := make(chan interface{})
	go func() {
		for d := range datas {
			fmt.Printf("id:%d ", d.job.id)
			fmt.Printf("%d - %d\n", d.job.i, d.square)
		}
		done <- true
	}()

	makeWP(workers)
	fmt.Printf(": %v\n", <-done)
}
