package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	port := ":8080"
	if len(os.Args) != 1 {
		port = os.Args[1]
	}

	s, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(data)) == "abc" {
			fmt.Println("exiting...")
			return
		}

		fmt.Print(">> ", string(data))
		t := time.Now()
		diyT := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(diyT))
	}
}
