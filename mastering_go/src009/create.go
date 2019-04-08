package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 10, "协程个数")
	flag.Parse()

	var wait sync.WaitGroup
	fmt.Printf("%#v\n", wait)
	for i := 0; i < *n; i++ {
		wait.Add(1)
		go func(x int) {
			defer wait.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	fmt.Printf("%#v\n", wait)
	wait.Wait()
	fmt.Printf("%#v\n", wait)

	fmt.Println()
}
