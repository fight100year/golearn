package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("look at me")
		}
	}()

	if len(os.Args) == 1 {
		panic("no more args")
	}

	fmt.Println("done")
}
