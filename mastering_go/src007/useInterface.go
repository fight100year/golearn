package main

import (
	"fmt"
	"math"

	"github.com/63isOK/golearn/diy"
)

type square struct {
	X float64
}

type circle struct {
	R float64
}

func (s square) Area() float64 {
	return s.X * s.X
}

func (s square) Perimeter() float64 {
	return 4 * s.X
}

func (c circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (c circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

func calc(x diy.Shape) {
	_, ok := x.(circle)
	if ok {
		fmt.Println("is a circle")
	}

	v, ok := x.(square)
	if ok {
		fmt.Println("is a square:", v)
	}

	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}

func main() {
	x := square{X: 10}
	fmt.Println("周长:", x.Perimeter())
	calc(x)

	y := circle{R: 5}
	calc(y)
}
