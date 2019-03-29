package main

type heapFloat32 []float32

func (n *heapFloat32) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	new := old[0 : len(old)-1]
	*n = new

	return x
}

func main() {
}
