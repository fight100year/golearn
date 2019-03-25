package main

import "fmt"

func return3(x int) (int, int, int) {
	return 2 * x, x * x, -x
}

func main() {
	fmt.Println(return3(10))
	n1, n2, n3 := return3(5)
	fmt.Println(n1, n2, n3)
}
