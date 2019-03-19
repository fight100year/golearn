package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func testMinMax() {
	if len(os.Args) == 1 {
		fmt.Println("please give me more floats number")
		os.Exit(1)
	}

	args := os.Args
	min, _ := strconv.ParseFloat(args[1], 64)
	max, _ := strconv.ParseFloat(args[1], 64)

	for i := 2; i < len(args); i++ {
		n, _ := strconv.ParseFloat(args[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("min:", min)
	fmt.Println("max:", max)
}

func main() {
	testMinMax()
	io.WriteString(os.Stdout, "hi")
	io.WriteString(os.Stderr, "hello")

}
