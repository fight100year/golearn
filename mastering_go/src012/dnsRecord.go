package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	domain := os.Args[1]
	nss, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, ns := range nss {
		fmt.Println(ns.Host)
	}

	mxs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, mx := range mxs {
		fmt.Println(mx.Host)
	}

}
