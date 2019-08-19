package chatterbox

import (
	"net"
	"strings"
	"sync"
)

type chatServer struct {
	net.Listener

	mu          sync.Mutex
	connections map[*client]net.Conn
	messages    chan message
}

// Collect keeps track of incoming connections, creates a client
// and labels the client with a room.
func (s *chatServer) Collect(room string) {
	for {
		conn, err := s.Accept()
		if err != nil {
			return
		}

		cl := newClient(conn)
		cl.room = room

		s.mu.Lock()
		s.connections[cl] = conn
		s.mu.Unlock()

		go cl.readFromConn(s.messages)
		cl.writeToConn(message{name: "admin", text: "Welcome, your handle is: " + cl.name})
	}
}

func (s *chatServer) drop(cl *client) {
	s.mu.Lock()
	delete(s.connections, cl)
	s.mu.Unlock()
}

func (s *chatServer) monitor() {
	for msg := range s.messages {

		if strings.HasPrefix(msg.text, "/") {
			f := strings.Fields(msg.text)
			switch f[0] {
			case "/close":
				s.connections[msg.from].Close()
				s.drop(msg.from)
				msg.text = "has left the server"

			case "/room":
				if len(f) != 2 {
					continue
				}
				msg.from.room = f[1]
				msg.from.writeToConn(message{name: "admin", text: "room changed to: " + f[1]})
				continue
			}
		}

		for client := range s.connections {
			if client.room == msg.room {
				client.writeToConn(msg)
			}
		}
	}
}

// NewServer returns a server acceptiong connections on `port` and begins monitoring
// for messages and commands.
func NewServer(port string) (*chatServer, error) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}

	serv := &chatServer{
		Listener:    listen,
		connections: make(map[*client]net.Conn),
		messages:    make(chan message),
	}

	go serv.monitor()
	return serv, nil
}
