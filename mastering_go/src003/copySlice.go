package main

import "fmt"

func main() {
	a6 := []int{1, 2, 3, 4, 5, 6}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	// copy(a6, a4)
	copy(a4, a6)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()

}
