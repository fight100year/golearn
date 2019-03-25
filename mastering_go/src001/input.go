package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f := os.Stdin
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() == "STOP" {
			fmt.Println("stop")

			return
		}

		fmt.Println("input >", s.Text())
	}
}
