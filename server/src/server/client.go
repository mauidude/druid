package main

import (
	"fmt"
	"io"
	"log"
	"sync/atomic"

	"golang.org/x/net/websocket"
)

const channelBufSize = 100

var maxID int = 0

// Chat client.
type Client struct {
	id     int
	ws     *websocket.Conn
	server *Server
	ch     chan *Message
	doneCh chan struct{}
}

// Create new chat client.
func NewClient(ws *websocket.Conn, server *Server) *Client {
	return &Client{
		id:     atomic.AddInt32(&maxID, 1),
		ws:     ws,
		server: server,
		ch:     make(chan *Message, channelBufSize),
		doneCh: make(chan struct{}),
	}
}

func (c *Client) Conn() *websocket.Conn {
	return c.ws
}

func (c *Client) Write(msg *Message) {
	select {
	case c.ch <- msg:
	default:
		c.server.Del(c)
		err := fmt.Errorf("client %d is disconnected.", c.id)
		c.server.Err(err)
	}
}

func (c *Client) Done() {
	c.doneCh <- struct{}{}
}

// Listen Write and Read request via chanel
func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

// Listen write request via chanel
func (c *Client) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {

		// send message to the client
		case msg := <-c.ch:
			log.Println("Send:", msg)
			websocket.JSON.Send(c.ws, msg)

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- struct{}{} // for listenRead method
			return
		}
	}
}

// Listen read request via chanel
func (c *Client) listenRead() {
	log.Println("Listening read from client")
	for {
		select {

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- struct{}{} // for listenWrite method
			return

		// read data from websocket connection
		default:
			var msg Message
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF {
				c.doneCh <- struct{}{}
			} else if err != nil {
				c.server.Err(err)
			} else {
				c.server.SendAll(&msg)
			}
		}
	}
}
