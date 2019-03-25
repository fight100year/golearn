package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	ver := runtime.Version()
	major := strings.Split(ver, ".")[0][2]
	minor := strings.Split(ver, ".")[1]
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)

	if m1 == 1 && m2 < 8 {
		fmt.Println("please use go version1.8 or higher")
		return
	}

	fmt.Println("version:", m1, m2)
}
