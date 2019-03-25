package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [...]int{0, 1, 2, 3, 4}
	point := &array[0]
	addr := uintptr(unsafe.Pointer(point))

	for i := 0; i < len(array)+2; i++ {
		point = (*int)(unsafe.Pointer(addr))
		fmt.Print(*point, " ")
		addr = uintptr(unsafe.Pointer(point)) + unsafe.Sizeof(array[0])
	}

	fmt.Println()
	point = (*int)(unsafe.Pointer(addr))
	fmt.Print("one more: ", *point, " ")
	fmt.Println()
	addr = uintptr(unsafe.Pointer(point)) + unsafe.Sizeof(array[0])
	point = (*int)(unsafe.Pointer(addr))
	fmt.Print("one more: ", *point, " ")

}
