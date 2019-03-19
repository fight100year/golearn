package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("please give me more floats")
		os.Exit(1)
	}

	args := os.Args
	err := errors.New("an error")
	k := 1
	var f float64

	for err != nil {
		if k >= len(args) {
			fmt.Println("none of the args is a float")
			return
		}
		f, err = strconv.ParseFloat(args[k], 64)
		k++
	}

	min, max := f, f
	for i := 2; i < len(args); i++ {
		f, err := strconv.ParseFloat(args[i], 64)
		if err == nil {
			if f < min {
				min = f
			}
			if f > max {
				max = f
			}
		}
	}

	fmt.Println("min:", min)
	fmt.Println("max:", max)
}
