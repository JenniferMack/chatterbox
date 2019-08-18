package chatterbox

import (
	"bufio"
	"net"
	"sync"
)

type chatroom struct {
	mu      sync.Mutex
	clients map[net.Conn]*client

	name     string
	messages chan []byte
}

func (cr *chatroom) Connect(c net.Conn) {
	client := &client{
		write:        bufio.NewWriter(c),
		roomMessages: cr.messages,
	}

	cr.mu.Lock()
	cr.clients[c] = client
	cr.mu.Unlock()
	go client.relay()
}

func NewRoom(name string) *chatroom {
	room := &chatroom{
		name:     name,
		messages: make(chan []byte),
		clients:  make(map[net.Conn]*client),
	}

	return room
}
