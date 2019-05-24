package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cnts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		cnts[input.Text()]++
	}

	for k, v := range cnts {
		fmt.Printf("%d\t%s\n", v, k)
	}
}
