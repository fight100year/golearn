package main

import (
	"fmt"
	"strings"
)

func main() {
	pos := strings.LastIndex("BKWRK 2 5", " ")
	pos1 := strings.Index("BKWRK 2 5", " ")
	fmt.Println(pos, pos1)
}
