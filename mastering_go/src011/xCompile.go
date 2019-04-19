package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("编译器：", runtime.Compiler)
	fmt.Println("os：", runtime.GOOS)
	fmt.Println("arch：", runtime.GOARCH)
	fmt.Println("go 版本：", runtime.Version())
}
