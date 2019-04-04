package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("vim-go")

	k := flag.Bool("k", true, "k")
	o := flag.Int("O", 1, "O")
	flag.Parse()

	valueK := *k
	valueO := *o

	fmt.Println("-k:", valueK)
	fmt.Println("-O:", valueO)
}
