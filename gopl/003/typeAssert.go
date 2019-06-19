package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Reader = os.Stdout
	fmt.Printf("%[1]v %[1]T\n", w)
	if w, ok := w.(io.ReadWriter); ok {
		fmt.Printf("%[1]v %[1]T\n", w)
	}
	if w, ok := w.(*bufio.Reader); ok {
		fmt.Printf("%[1]v %[1]T\n", w)
	} else {
		fmt.Println("no")
	}
}
