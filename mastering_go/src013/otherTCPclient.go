package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	serverInfo := "localhost:8080"
	if len(os.Args) != 1 {
		serverInfo = os.Args[1]
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverInfo)
	if err != nil {
		fmt.Println("ResolveTCPAddr:", err.Error())
	}

	c, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		fmt.Println("DialTCP:", err.Error())
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		txt, _ := reader.ReadString('\n')
		fmt.Fprintf(c, txt+"\n")

		msg, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("<< " + msg)
		if strings.TrimSpace(string(txt)) == "stop" {
			fmt.Println("tcp client exiting...")
			return
		}
	}
}
