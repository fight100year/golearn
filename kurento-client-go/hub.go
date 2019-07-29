package main

import (
	"fmt"
)

type hub struct {
	clients    map[*client]bool
	message    chan *clientMessage
	register   chan *client
	unregister chan *client
}

func newHub() *hub {
	return &hub{
		clients:    make(map[*client]bool),
		message:    make(chan *clientMessage),
		register:   make(chan *client),
		unregister: make(chan *client),
	}
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true
		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.send)
			}
		case message := <-h.message:
			fmt.Println(string(message.message))
			message.client.send <- message.message
		case host := <-retry:
			reconnectKMS(h, host)
		}
	}
}
