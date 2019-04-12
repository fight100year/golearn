package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func set(i int) {
	writeValue <- i
}

func read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case i := <-writeValue:
			value = i
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("miss param")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	rand.Seed(time.Now().Unix())
	go monitor()

	var w sync.WaitGroup
	w.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer w.Done()
			set(rand.Intn(10 * n))
		}()
	}

	w.Wait()
	fmt.Println(read())
}
