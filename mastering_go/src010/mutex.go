package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	// m sync.Mutex
	v int
)

func change(i int) {
	// m.Lock()
	time.Sleep(time.Second)
	v++
	// m.Unlock()
}

func read() int {
	// m.Lock()
	a := v
	// m.Unlock()
	return a
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var w sync.WaitGroup

	fmt.Printf("%d ", read())
	for i := 0; i < num; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			change(i)
			fmt.Printf(" -> %d", read())
		}(i)
	}

	w.Wait()
	fmt.Printf("-> %d\n", read())
}
