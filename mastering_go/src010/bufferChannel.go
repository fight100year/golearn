package main

import (
	"fmt"
)

func main() {
	num := make(chan int, 5)
	cnt := 10

	for i := 0; i < cnt; i++ {
		select {
		case num <- i:
		default:
			fmt.Println("not enouth space for", i)
		}
	}

	for i := 0; i < cnt+5; i++ {
		select {
		case n := <-num:
			fmt.Println(n)
		default:
			fmt.Println("not enouth space for", i)
			break
		}
	}
}
