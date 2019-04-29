package main

import (
	"fmt"
	"math"
	"net"
	"net/rpc"
	"os"
	"sharerpc"
)

// InterfaceF rpc服务类型
type InterfaceF struct{}

// Multiply 实现的rpc之一 乘积
func (i *InterfaceF) Multiply(args *sharerpc.TypeF, reply *float64) error {
	*reply = args.F1 * args.F2
	return nil
}

// Power 实现的rpc之一 幂
func (i *InterfaceF) Power(args *sharerpc.TypeF, reply *float64) error {
	*reply = power(args.F1, args.F2)
	return nil
}

func power(x, y float64) float64 {
	return math.Pow(x, y)
}

func main() {
	port := ":8888"
	if len(os.Args) != 1 {
		port = os.Args[1]
	}

	rpc.Register(new(InterfaceF))
	addr, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}
}
