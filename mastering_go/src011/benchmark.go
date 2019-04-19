package main

import "fmt"

func f1(n int) int {
	if n == 0 {
		return n
	}
	if n == 1 {
		return n
	}

	return f1(n-1) + f1(n-2)
}

func f2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return f2(n-1) + f2(n-2)
}

func f3(n int) int {
	m := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i < 2 {
			f = 1
		} else {
			f = m[i-1] + m[i-2]
		}
		m[i] = f
	}

	return m[n]
}

func main() {
	fmt.Println(f1(30))
	fmt.Println(f2(30))
	fmt.Println(f3(30))
}
