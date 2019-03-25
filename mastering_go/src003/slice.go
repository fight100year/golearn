package main

import (
	"fmt"
	"unsafe"
)

func reSlice() {
	s := make([]int, 5)
	rs := s[1:3]
	fmt.Println(s)
	fmt.Println(rs)

	rs[0] = 123
	rs[1] = 456
	fmt.Println(s)
	fmt.Println(rs)

	fmt.Println()
	fmt.Println()
}

func printSlice(x []int) {
	for _, n := range x {
		fmt.Print(n, " ")
	}
	fmt.Println()
}

func lenCap() {
	fmt.Println("--------> 初始信息")
	s := []int{1, 2, 3}
	rs := s[2:]
	fmt.Printf("s: ")
	printSlice(s)
	fmt.Printf("s cap: %d, len: %d\n", cap(s), len(s))
	fmt.Printf("rs: ")
	printSlice(rs)
	fmt.Printf("rs cap: %d, len: %d\n", cap(rs), len(rs))

	// rs = append(rs, 4)
	// fmt.Printf("s: ")
	// printSlice(s)
	// fmt.Printf("s cap: %d, len: %d\n", cap(s), len(s))
	// fmt.Printf("rs: ")
	// printSlice(rs)
	// fmt.Printf("rs cap: %d, len: %d\n", cap(rs), len(rs))

	fmt.Println("--------> 切片追加")
	//rs[0] = 100
	s = append(s, 100)
	fmt.Printf("s: ")
	printSlice(s)
	fmt.Printf("s cap: %d, len: %d\n", cap(s), len(s))
	fmt.Printf("rs: ")
	printSlice(rs)
	fmt.Printf("rs cap: %d, len: %d\n", cap(rs), len(rs))

	fmt.Println("--------> 切片追加")
	rs[0] = 4
	fmt.Printf("s: ")
	printSlice(s)
	fmt.Printf("s cap: %d, len: %d\n", cap(s), len(s))
	fmt.Printf("rs: ")
	printSlice(rs)
	fmt.Printf("rs cap: %d, len: %d\n", cap(rs), len(rs))
}

func main() {
	reSlice()
	lenCap()

	fmt.Println("--------> 数组")
	a := [5]int{1, 2, 3, 4, 5}
	s := a[:]
	rs := s[3:]

	fmt.Printf("s: ")
	printSlice(s)
	fmt.Printf("s cap: %d, len: %d\n", cap(s), len(s))
	fmt.Printf("rs: ")
	printSlice(rs)
	fmt.Printf("rs cap: %d, len: %d\n", cap(rs), len(rs))
	fmt.Println("a:", unsafe.Pointer(&a[0]), "s:", unsafe.Pointer(&s[0]), "rs:", unsafe.Pointer(&rs[0]))

	fmt.Println("--------> 修改原始切片")
	// s = append(s, 6)
	s[3] = 0
	fmt.Printf("s: ")
	printSlice(s)
	fmt.Printf("s cap: %d, len: %d\n", cap(s), len(s))
	fmt.Printf("rs: ")
	printSlice(rs)
	fmt.Printf("rs cap: %d, len: %d\n", cap(rs), len(rs))
	fmt.Println("a:", unsafe.Pointer(&a[0]), "s:", unsafe.Pointer(&s[0]), "rs:", unsafe.Pointer(&rs[0]))
	fmt.Println("a:", a)

	rs = nil
	fmt.Println("rs=nil:", rs)
}
