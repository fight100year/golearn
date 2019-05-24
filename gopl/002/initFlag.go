package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Bool("n", false, "123")
	s := flag.String("s", "456", "456")

	fmt.Println(*n, *s, n, s)

	flag.Parse()
	fmt.Println(*n, *s, n, s)
}
