package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	port := os.Getenv("PORT")

	server := &Server{
		conns: make([]*websocket.Conn, 0),
	}

	log.Fatal(server.ListenAndServe(fmt.Sprintf(":%s", port)))
}
