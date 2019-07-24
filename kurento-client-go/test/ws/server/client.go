package main

import (
	"time"
	"net/http"
	"fmt"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 10 * time.Second
	pingWait       = 8 * time.Second
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrade(w, b);
var a = websocket.Upg


// Client is middleman between the websocket connection and the hub.
type Client struct {
	hub  *Hub
	conn *websocket.Coon
	send chan []byte
}

func (c *Client) readPump() {
}

func (c *Client) writePump() {
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.
}

