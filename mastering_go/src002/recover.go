package main

import "fmt"

func a() {
	fmt.Println("inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("recover() in a()")
		}
	}()
	fmt.Println("call b()")
	b()
	fmt.Println("call b() done")
	fmt.Println("a() done")
}

func b() {
	fmt.Println("inside b()")
	panic("panic in b()")
	fmt.Println("b() done")
}

func main() {
	a()
	fmt.Println("main() done")
}
