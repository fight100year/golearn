package main

import (
	"fmt"
	"sort"
)

type diy struct {
	name string
	x    int
	y    int
}

func main() {
	s := make([]diy, 0)
	s = append(s, diy{"dd", 7, 8})
	s = append(s, diy{"ee", 9, 10})
	s = append(s, diy{"aa", 1, 2})
	s = append(s, diy{"bb", 3, 7})
	s = append(s, diy{"cc", 5, 6})
	fmt.Println("0:", s)

	sort.SliceStable(s, func(i, j int) bool {
		return s[i].x < s[j].x
	})
	fmt.Println("x:", s)

	sort.SliceStable(s, func(i, j int) bool {
		return s[i].y < s[j].y
	})
	fmt.Println("y:", s)
}
