package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOMAXRPOCS: ", runtime.GOMAXPROCS(0))
}
