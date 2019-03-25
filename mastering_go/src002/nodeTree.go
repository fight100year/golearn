package main

import "fmt"

func abc(i int) {
	b := true
	fmt.Println(i, b)
}

func main() {
	aa := 1
	bb := 1

	abc(aa)
	abc(bb)
}
