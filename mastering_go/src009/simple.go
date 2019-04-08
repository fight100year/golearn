package main

import (
	"fmt"
	"time"
)

func fun() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}

func main() {
	go fun()
	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()

	time.Sleep(1 * time.Second)
}
