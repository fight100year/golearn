package main

import (
	"fmt"
	"net/http"
	"time"
)

func page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "ws need get", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func wsServe(hub *hub) {
	http.HandleFunc("/", page)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	go func() {
		err := http.ListenAndServe(":9000", nil)
		if err != nil {
			fmt.Println(err)
		}
	}()
}

// wsClient 处理kms客户端的连接
// 函数内部可扩展,用以使用多个kms服务
func wsClient(hub *hub) {
	go clientWs(hub, "192.168.10.180:8766")
}

func main() {
	hub := newHub()
	go hub.run()

	wsServe(hub)
	wsClient(hub)

	time.Sleep(10 * time.Minute)
}
