package main

import (
	"fmt"
	"net/http"
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

func wsServe(hub *Hub) {
	http.HandleFunc("/", page)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func wsClient(hub *Hub) {

}

func main() {
	hub := newHub()
	go hub.run()

	wsServe(hub)
	wsClient(hub)

}
