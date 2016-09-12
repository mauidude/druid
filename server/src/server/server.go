package main

import (
	"log"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

type Server struct {
	mutex     sync.RWMutex
	messages  []*Message
	clients   map[int]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan *Message
	doneCh    chan struct{}
	errCh     chan error
}

// Create new chat server.
func NewServer() *Server {
	messages := []*Message{}
	clients := make(map[int]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	sendAllCh := make(chan *Message)
	doneCh := make(chan struct{})
	errCh := make(chan error)

	return &Server{
		messages,
		clients,
		addCh,
		delCh,
		sendAllCh,
		doneCh,
		errCh,
	}
}

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) Del(c *Client) {
	s.delCh <- c
}

func (s *Server) SendAll(msg *Message) {
	s.sendAllCh <- msg
}

func (s *Server) Done() {
	s.doneCh <- struct{}{}
}

func (s *Server) Err(err error) {
	s.errCh <- err
}

func (s *Server) sendAll(msg *Message) {
	s.mutex.RLock()
	defer s.mutex.Unlock()

	for _, c := range s.clients {
		c.Write(msg)
	}
}

// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen(path string) {

	log.Println("Listening server...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		client := NewClient(ws, s)
		s.Add(client)
		client.Listen()
	}
	http.Handle(path, websocket.Handler(onConnected))
	log.Println("Created handler")

	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			log.Println("Added new client")
			s.mutex.Lock()
			s.clients[c.id] = c

			count := len(s.clients)
			s.mutex.Unlock()

			log.Println("Now", count, "clients connected.")

		// del a client
		case c := <-s.delCh:
			log.Println("Delete client")

			s.mutex.Lock()
			delete(s.clients, c.id)
			s.mutex.Unlock()

		// broadcast message for all clients
		case msg := <-s.sendAllCh:
			log.Println("Sending all")
			s.sendAll(msg)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}
