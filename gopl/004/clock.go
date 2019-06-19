package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "172.17.0.2:8000")
	if err != nil {
		log.Fatal(err)

		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05.00000\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
