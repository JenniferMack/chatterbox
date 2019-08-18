package chatterbox

import (
	"bufio"
	"math/rand"
	"net"
	"sync"
	"time"
)

type chatroom struct {
	mu      sync.Mutex
	clients map[net.Conn]*client

	name     string
	messages chan []byte
}

func (cr *chatroom) Close() {
	cr.Send([]byte("server closing\n"))
	time.Sleep(100 * time.Millisecond)

	for conn, client := range cr.clients {
		close(client.recv)
		conn.Close()
	}
}

func (cr *chatroom) Accept(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		cr.Connect(conn, time.Now().UnixNano())
	}
}

func (cr *chatroom) Connect(c net.Conn, seed int64) {
	client := &client{
		name:  makeName(seed),
		write: bufio.NewWriter(c),
		read:  bufio.NewReader(c),
		send:  cr.messages,
		recv:  make(chan []byte),
	}

	cr.mu.Lock()
	cr.clients[c] = client
	cr.mu.Unlock()
	go client.relay()
	go client.monitor()

	cr.Send([]byte("#> " + client.name + " has entered " + cr.name + "\n"))
}

func (cr *chatroom) Send(msg []byte) {
	for _, client := range cr.clients {
		client.recv <- msg
	}
}

func (cr *chatroom) listen() {
	for byt := range cr.messages {
		cr.Send(byt)
	}
}

func NewRoom(name string) *chatroom {
	room := &chatroom{
		name:     name,
		messages: make(chan []byte),
		clients:  make(map[net.Conn]*client),
	}
	go room.listen()

	return room
}

func makeName(seed int64) string {
	fn := []string{
		"alpha", "bravo", "charlie", "delta", "echo", "foxtrot",
		"golf", "hotel", "india", "juliet", "kilo", "lima", "mike",
		"november", "oscar", "papa", "quebec", "romeo", "sierra", "tango",
		"uniform", "victor", "whiskey", "xray", "yankee", "zulu",
	}
	rn := rand.New(rand.NewSource(seed))
	p1 := rn.Intn(len(fn) - 1)
	p2 := rn.Intn(len(fn) - 1)
	return fn[p1] + "_" + fn[p2]
}
