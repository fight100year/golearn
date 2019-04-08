package main

import (
	"fmt"
	"time"
)

func write(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
	fmt.Println(<-c)
	close(c)
	fmt.Println(x)
}

func main() {
	c := make(chan int)

	go write(c, 10)
	time.Sleep(time.Second)

	fmt.Println("read:", <-c)
	time.Sleep(time.Second)

	_, ok := <-c
	if ok {
		fmt.Println("channel is open")
	} else {
		fmt.Println("channel is close")
	}
}
