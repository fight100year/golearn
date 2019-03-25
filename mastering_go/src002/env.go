package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Compiler)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.Version())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}
