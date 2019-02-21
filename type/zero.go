package main

import (
	"fmt"
	"github.com/63isOK/golearn/swap"
)

func main(){
	var (
		i int
		f float64
		b bool
		s string
	)

	fmt.Println("%v %v %v %q \n", i, f, b, s)

	x, y := swap.Swap("123", "abc")
	fmt.Println("123-abc, swap: ", x, y);
}
