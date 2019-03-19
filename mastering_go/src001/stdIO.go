package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func print() {
	str := ""
	args := os.Args
	if len(args) == 1 {
		str = "please give me one argument"
	} else {
		str = args[1]
	}

	io.WriteString(os.Stdout, str)
	io.WriteString(os.Stdout, "\n")
}

func read() {
	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}

func main() {
	print()
	read()
}
