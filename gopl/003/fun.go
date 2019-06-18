package main

import (
	"fmt"
)

func squares() func() int {
	x := 0
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(squares()())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(squares()())
	fmt.Println(squares()())
	fmt.Println(f())
	fmt.Println(f())
}
