package main

import (
	"fmt"
	"time"
)

func main() {
	nums := make(chan int)
	sqra := make(chan int)

	go func() {
		for x := 0; x < 30; x++ {
			nums <- x
			time.Sleep(100 * time.Millisecond)
		}
		close(nums)
	}()

	go func() {
		for x := range nums {
			sqra <- x * x
		}
		close(sqra)
	}()

	for {
		fmt.Println(<-sqra)
	}
}
