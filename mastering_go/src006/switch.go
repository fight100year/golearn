package main

import "fmt"

type s struct {
	x float64
}

type c struct {
	r float64
}

type r struct {
	x float64
	y float64
}

func print(x interface{}) {
	switch v := x.(type) {
	case s:
		fmt.Println("this is s")
	case c:
		fmt.Println("this is c")
	case r:
		fmt.Println("this is r")
	default:
		fmt.Printf("unknow type: %T\n", v)
	}
}

func main() {
	x := c{r: 10}
	print(x)

	y := s{x: 10}
	print(y)

	z := r{x: 1, y: 2}
	print(z)

	print(1)
}
