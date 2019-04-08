package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var closeA = false
var data = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if closeA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Println(x, " ")
		_, ok := data[x]
		if ok {
			closeA = true
		} else {
			data[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	sum := 0
	for x := range in {
		sum += x
	}

	fmt.Println("sum:", sum)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: xx min max")
		os.Exit(1)
	}

	min, _ := strconv.Atoi(os.Args[1])
	max, _ := strconv.Atoi(os.Args[2])
	if min > max {
		return
	}

	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)

	go first(min, max, A)
	go second(B, A)
	third(B)
}
