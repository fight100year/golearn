package main

import (
	"fmt"
)

func f11() {
	fmt.Println("f1()")
	return
	fmt.Println("here")
}

func f22() {
	fmt.Println("f2()")
}

func main() {
	fmt.Println("done")
}
