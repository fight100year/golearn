package main

import (
	"fmt"
)

func f1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
	fmt.Println()
}

// 这种写法，输出0 0 0
// defer 函数只有参数是立马计算
// f2的写法，如果不是故意这么写，其他情况会引起bug
func f2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

// 最推荐的写法
func f3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
	fmt.Println()
}

func main() {
	f1()
	f2()
	f3()
}
