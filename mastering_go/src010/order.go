package main

import (
	"fmt"
	"time"
)

func funa(a, b chan struct{}) {
	<-a
	fmt.Println("A()")
	time.Sleep(time.Second)
	close(b)
}

func funb(a, b chan struct{}) {
	<-a
	fmt.Println("B()")
	close(b)
}

func fund(a chan struct{}) {
	<-a
	fmt.Println("C()")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	// 等待z
	go fund(z)

	// 等待x，关闭y
	go funa(x, y)

	go fund(z)

	// 等待y，关闭z
	go funb(y, z)

	go fund(z)

	close(x)
	time.Sleep(3 * time.Second)
}
