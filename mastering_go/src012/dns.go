package main

import (
	"fmt"
	"net"
	"os"
)

func lookIP(address string) ([]string, error) {
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func lookHostname(hostname string) ([]string, error) {
	ips, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("need params")
		return
	}

	ip := net.ParseIP(os.Args[1])
	if ip == nil {
		ips, err := lookHostname(os.Args[1])
		if err != nil {
			for _, x := range ips {
				fmt.Println(x)
			}
		}
	} else {
		hosts, err := lookIP(os.Args[1])
		if err == nil {
			for _, x := range hosts {
				fmt.Println(x)
			}
		}
	}
}
