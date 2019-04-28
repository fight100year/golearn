package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("interface name: %v\n", i.Name)
		fmt.Println("interface flags", i.Flags)
		fmt.Println("interface mtu", i.MTU)
		fmt.Println("interface addr", i.HardwareAddr)
		fmt.Println()
	}
}
