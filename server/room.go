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
	clients map[*client]net.Conn

	name     string
	messages chan []byte
}

// Close ends client connections and closes channels.
func (cr *chatroom) Close() {
	cr.send([]byte("server closing\n"))
	time.Sleep(100 * time.Millisecond)

	for client, conn := range cr.clients {
		close(client.recv)
		conn.Close()
	}
	close(cr.messages)
}

// Accept monitors the server and accepts incoming connections.
func (cr *chatroom) Accept(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		cr.connect(conn, time.Now().UnixNano())
	}
}

func (cr *chatroom) connect(c net.Conn, seed int64) {
	client := &client{
		name:  makeName(seed),
		write: bufio.NewWriter(c),
		read:  bufio.NewReader(c),
		send:  cr.messages,
		recv:  make(chan []byte),
	}

	cr.mu.Lock()
	cr.clients[client] = c
	cr.mu.Unlock()
	go client.relay()
	go client.monitor(cr.drop)

	cr.send([]byte("#> " + client.name + " has entered " + cr.name + "\n"))
}

func (cr *chatroom) drop(cl *client) {
	cr.mu.Lock()
	cr.clients[cl].Close()
	delete(cr.clients, cl)
	cr.mu.Unlock()
}

func (cr *chatroom) send(msg []byte) {
	for client := range cr.clients {
		client.recv <- msg
	}
}

func (cr *chatroom) listen() {
	for byt := range cr.messages {
		cr.send(byt)
	}
}

// NewRoom creats a chatroom and starts a listening thread.
func NewRoom(name string) *chatroom {
	room := &chatroom{
		name:     name,
		messages: make(chan []byte),
		clients:  make(map[*client]net.Conn),
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
