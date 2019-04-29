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

	c, err := net.Dial("tcp", serverInfo)
	if err != nil {
		fmt.Println(err)
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
