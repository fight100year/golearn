package main

import (
	"fmt"
	"github.com/mactsouk/go/simpleGitHub"
)

func testGeneral() {
	i := 123
	s := "abc"
	b := true

	fmt.Printf("%t\n", b)
	fmt.Printf("%v,%#v,%T, %%\n", s, s, s)
	fmt.Printf("%v,%#v,%T, %%\n", i, i, i)
}

func testInteger() {
	i := 97
	fmt.Printf("%b, %c, %d, %o, %q, %x, %X, %U\n", i, i, i, i, i, i, i, i)
}

func main() {
	fmt.Println(simpleGitHub.AddTwo(1, 2))

	testGeneral()
	testInteger()

}
