package main

import "fmt"

type a struct {
	XX int
	YY int
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

func (A a) A() {
	fmt.Println("a.A()")
}

func (B b) A() {
	fmt.Println("b.A()")
}

func sameFunc() {
	var i c
	i.A.A()
	i.B.A()
}

type first struct{}

func (a first) F() {
	a.shared()
}

func (a first) shared() {
	fmt.Println("first.shared()")
}

type second struct{ first }

func (a second) shared() {
	fmt.Println("second.shared()")
}

func main() {
	sameFunc()

	fmt.Println()

	first{}.F()
	second{}.shared()
	j := second{}
	i := j.first
	i.F()

}
