package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func f1(i int) {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c.Done():
		fmt.Println("f1():", c.Err())
		return
	case r := <-time.After(time.Duration(i) * time.Second):
		fmt.Println("f1():", r)
	}
}

func f2(i int) {
	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Duration(i)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c.Done():
		fmt.Println("f2():", c.Err())
		return
	case r := <-time.After(time.Duration(i) * time.Second):
		fmt.Println("f2():", r)
	}
}

func f3(i int) {
	c := context.Background()
	deadline := time.Now().Add(time.Duration(2*i) * time.Second)
	c, cancel := context.WithDeadline(c, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c.Done():
		fmt.Println("f3():", c.Err())
		return
	case r := <-time.After(time.Duration(i) * time.Second):
		fmt.Println("f3():", r)
	}
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	f1(i)
	f2(i)
	f3(i)
}
