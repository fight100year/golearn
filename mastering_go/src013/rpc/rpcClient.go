package main

import (
	"fmt"
	"net/rpc"
	"os"
	"sharerpc"
)

func main() {
	addr := ":8888"
	if len(os.Args) != 1 {
		addr = os.Args[1]
	}

	c, err := rpc.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	args := sharerpc.TypeF{F1: 3.0, F2: 2.0}
	var result float64
	err = c.Call("InterfaceF.Multiply", args, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("reply (multiply): %f\n", result)

	err = c.Call("InterfaceF.Power", args, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("reply (power): %f\n", result)
}
