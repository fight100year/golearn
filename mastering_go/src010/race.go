package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var w sync.WaitGroup
	var i int
	k := make(map[int]int)
	k[1] = 12
	w.Add(n)
	for i = 0; i < n; i++ {
		go func() {
			defer w.Done()
			k[i] = i
		}()
	}

	k[2] = 10
	w.Wait()
	fmt.Printf("k = %#v\n", k)
}
