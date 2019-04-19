// Package ex 是用来测试例子函数的
// 只是一个例子
package ex

// P 圆周率
const P = 3.14

// F1 计算斐波拉切数列
func F1(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return F1(n-1) + F1(n-2)
}

// S1 计算字符串长度jj
func S1(s string) int {
	return len(s)
}
