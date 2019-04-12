package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func f1(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}

	time.Sleep(time.Millisecond)
	return int64(f1(n-1)) + int64(n-2)
}

func f2(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}

	time.Sleep(50 * time.Millisecond)
	return fn[n]
}

// n1 计算是否是质数
func n1(n int) bool {
	k := math.Floor(float64(n/2 + 1))
	for i := 2; i < int(k); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func n2(n int) bool {
	for i := 2; i < n; i++ {
		if (n % i) == 0 {
			return false
		}
	}

	return true
}

func main() {
	cpu, err := os.Create("cpu.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cpu.Close()

	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()

	sum := 0
	for i := 2; i < 100000; i++ {
		if n1(i) {
			sum++
		}
	}
	fmt.Println("(方法1)100000内的质数数量有", sum)

	sum = 0
	for i := 2; i < 100000; i++ {
		if n2(i) {
			sum++
		}
	}
	fmt.Println("(方法2)100000内的质数数量有", sum)

	fmt.Println("递归计算：")
	for i := 1; i < 90; i++ {
		n := f1(i)
		fmt.Print(n, " ")
	}
	fmt.Println()
	fmt.Println("队列计算：")
	for i := 1; i < 90; i++ {
		n := f2(i)
		fmt.Print(n, " ")
	}
	fmt.Println()
	runtime.GC()

	// memory profile
	memory, err := os.Create("memory.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer memory.Close()

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("operation failed")
		}
		time.Sleep(50 * time.Millisecond)
	}
	err = pprof.WriteHeapProfile(memory)
	if err != nil {
		fmt.Println(err)
		return
	}
}
