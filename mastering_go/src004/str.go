package main

import (
	"fmt"
)

func main() {
	s := "1234aafwef"

	fmt.Println("len:", len(s))

	for _, x := range s {
		fmt.Printf("%x ", x)
	}

	fmt.Println()

	fmt.Printf("q: %q\n", s)
	fmt.Printf("+q: %+q\n", s)
	fmt.Printf(" x: % x\n", s)
}
