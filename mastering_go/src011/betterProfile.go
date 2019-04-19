package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func n1(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()

	cnt := 0
	for i := 2; i < 20000; i++ {
		if n1(i) {
			cnt++
		}
	}
	fmt.Println("cnt:", cnt)
}
