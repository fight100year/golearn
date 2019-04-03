package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	var i interface{} = 123
	k, ok := i.(int)
	if ok {
		fmt.Println("success:", k)
	}

	v, ok := i.(float64)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("failed")
	}

	a := i.(int)
	fmt.Println("no check:", a)
	b := i.(bool)
	fmt.Println(b)
}
