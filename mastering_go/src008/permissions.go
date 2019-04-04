package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]
	info, _ := os.Stat(arg)
	mode := info.Mode()
	fmt.Println(arg, "mode is", mode.String()[1:10])
}
