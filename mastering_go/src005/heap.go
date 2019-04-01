package main

import (
	"container/heap"
	"fmt"
)

type heapFloat32 []float32

func (n *heapFloat32) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	new := old[0 : len(old)-1]
	*n = new

	return x
}

func (n *heapFloat32) Push(x interface{}) {
	*n = append(*n, x.(float32))
}

func (n heapFloat32) Len() int {
	return len(n)
}

func (n heapFloat32) Less(a, b int) bool {
	return n[a] < n[b]
}

func (n heapFloat32) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

func main() {
	h := &heapFloat32{6.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println("%v\n", h)
	heap.Init(h)
	fmt.Println("%v\n", h)

	h.Push(float32(12.0))
	h.Push(float32(-12.0))
	fmt.Println("%v\n", h)

	heap.Init(h)
	fmt.Println("%v\n", h)
}
