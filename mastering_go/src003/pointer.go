package main

import "fmt"

func main() {
	i := 1
	p := &i

	fmt.Println("memory:", p)
	fmt.Println("value:", *p)
}
