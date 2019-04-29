package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func f(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

func diyHandler(c net.Conn) {
	for {
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		temp := strings.TrimSpace(string(data))
		if temp == "stop" {
			break
		}

		fibo := "-1\n"
		n, err := strconv.Atoi(temp)
		if err == nil {
			fibo = strconv.Itoa(f(n)) + "\n"
		}
		c.Write([]byte(string(fibo)))
	}

	time.Sleep(5 * time.Second)
	c.Close()
}

func main() {
	port := ":8080"
	if len(os.Args) != 1 {
		port = os.Args[1]
	}

	l, err := net.Listen("tcp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go diyHandler(c)
	}
}
