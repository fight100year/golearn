package main

import "fmt"

func testBreak() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("i:", i)

		if i < 3 {
			continue
		}

		if i < 7 {
			fmt.Println(i)
		}

		if i > 7 {
			break
		}
	}
}

func testLoop() {
	i := 5
	for {
		if i--; i < 0 {
			break
		}
		fmt.Println(i)
	}
}

func testRange() {
	array := [5]int{1, 2, 3, 4, 5}
	for i, value := range array {
		fmt.Println(i, value)
	}
}

func main() {
	testBreak()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	testLoop()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	testRange()
}
