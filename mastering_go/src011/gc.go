package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func p(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)

	fmt.Println("mem.Alloc", mem.Alloc)
	fmt.Println("mem.TotalAlloc", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc", mem.HeapAlloc)
	fmt.Println("mem.NumGC", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()

	var m runtime.MemStats
	p(m)
	for i := 0; i < 3; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("failed")
		}
	}
	p(m)
	for i := 0; i < 15; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("failed")
		}
		time.Sleep(time.Millisecond)
	}
	p(m)
}
