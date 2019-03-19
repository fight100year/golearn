package main

import (
	_ "errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	sum := 0.0
	num := 0

	for i := 1; i < len(args); i++ {
		f, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			num++
			sum += f
		}
	}

	if num == 0 {
		fmt.Println("please give me some valid number")
		os.Exit(1)
	}

	fmt.Println("sum:", sum)
	fmt.Println("average:", sum/float64(num))
}
