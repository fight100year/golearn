package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var (
	retry chan string = make(chan string)
)

// func init() {
//     retry = make(chan string)
// }

func clientWs(hub *hub, host string) {
	if len(host) == 0 {
		fmt.Println("kms host is invalid:", host)
		return
	}

	u := url.URL{Scheme: "ws", Host: host, Path: "/kurento"}
	fmt.Printf("connect to kms: %s\n", host)

	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 3 * time.Second,
	}
	conn, _, err := dialer.Dial(u.String(), nil)

	if err != nil {
		fmt.Println(host, "- dial:", err)
		retry <- host

		return
	}

	c := &client{hub: hub, conn: conn, send: make(chan []byte, 256), host: host}
	c.hub.register <- c

	go c.writePump()
	go c.readPump()
}

func reconnectKMS(hub *hub, host string) {
	go func() {
		time.Sleep(3 * time.Second)
		clientWs(hub, host)
	}()
}
