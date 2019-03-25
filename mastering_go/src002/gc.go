package main

import (
	"fmt"
	"runtime"
	_ "time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("---")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 500)
		if s == nil {
			fmt.Println("operation failed")
		}
		// time.Sleep(50 * time.Millisecond)
	}

	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 10000000)
		if s == nil {
			fmt.Println("operation failed")
		}
		// time.Sleep(1 * time.Second)
	}

	printStats(mem)
}
