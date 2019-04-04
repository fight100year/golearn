package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func cat(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}
	return nil
}

func main() {
	file := ""
	arg := os.Args
	if len(arg) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for i := 0; i < len(arg); i++ {
		file = arg[i]
		err := cat(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
