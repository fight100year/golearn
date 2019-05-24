package main

import "fmt"

func main() {
	a := 123
	b := "hello"
	fmt.Printf("%[1]x %[1]o %[1]T, %[2]T, %[2]v\n", a, b)

	c := "a"
	d := "分"
	fmt.Println(len(c), len(d))
	for _, x := range c {
		fmt.Printf("a: %T\n", x)
	}
	for _, x := range c {
		fmt.Printf("d: %T\n", x)
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("c: %T\t", c[i])
	}
	for i := 0; i < len(d); i++ {
		fmt.Printf("d: %T\t", d[i])
	}

	fmt.Println()
	var e1 float32 = 3210.1234567890123456789
	var e2 float32 = 210.1234567890123456789
	var e3 float32 = 10.1234567890123456789
	var e4 float32 = 0.1234567890123456789
	var f = 3210.1234567890123456789
	fmt.Println(e1, e2, e3, e4, f)

	s1 := `abc
	123
	456
	...`
	fmt.Println(s1)

	fmt.Printf("%    d\n", 123)
}
