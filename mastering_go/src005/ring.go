package main

import (
	"container/ring"
	"fmt"
)

var size = 5

func main() {
	r := ring.New(size)
	fmt.Println("empty ring:", *r)

	for i := 0; i < r.Len()-1; i++ {
		r.Value = i
		r = r.Next()
	}

	r.Value = 2

	sum := 0
	r.Do(func(x interface{}) {
		t := x.(int)
		sum += t
	})
	fmt.Println("sum:", sum)

	for i := 0; i < r.Len()*3; i++ {
		r = r.Next()
		fmt.Print(r.Value, " ")
	}
}
