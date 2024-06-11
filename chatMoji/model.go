package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	join       chan *Client
	disconnect chan *Client
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		join:       make(chan *Client),
		disconnect: make(chan *Client),
	}
}
func (h *Hub) run() {
	for {
		select {
		case client := <-h.join:
			h.clients[client] = true
		case client := <-h.disconnect:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}

		}
	}
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

const (
	writeTime  = 20 * time.Second
	pongPeriod = 10 * time.Second
	pingPeriod = 8 * pongPeriod / 10 // shorter then pong
	readLimit  = 512
)

var (
	newline = []byte{'\n'}
)

func (c *Client) writeContent() {
	pingTicker := time.NewTicker(pingPeriod)
	c.conn.SetWriteDeadline(time.Time{}) /*never close*/ // time.Now().Add(writeTime) will close it after some time
	defer func() {
		pingTicker.Stop()
		c.conn.Close()
		log.Println("Closing conn of client", c)
	}()
	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println("Error on creating a next writer", err)
				return
			}
			w.Write(msg)
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}
			if err := w.Close(); err != nil {
				log.Println("Error on closing writer", err)
				return
			}
		case <-pingTicker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("Ping message failed on client", c)
				return
			}
		}
	}
}
func (c *Client) readContent() {
	defer func() {
		c.hub.disconnect <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(readLimit)
	c.conn.SetReadDeadline(time.Now().Add(pongPeriod))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongPeriod))
		return nil
	})
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("unexpected close", err)
			}
			return
		}
		message = bytes.Trim(message, "\n")
		c.hub.broadcast <- message
	}
}

func serveWS(h *Hub, w http.ResponseWriter, r *http.Request) {
	updater := websocket.Upgrader{ReadBufferSize: 512, WriteBufferSize: 512}
	conn, err := updater.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during update", err)
		return
	}
	client := &Client{hub: h, conn: conn, send: make(chan []byte, 1)}
	client.hub.join <- client

	go client.writeContent()
	go client.readContent()

}
