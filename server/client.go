package chatterbox

import (
	"bufio"
	"math/rand"
	"net"
	"strings"
	"time"
)

type client struct {
	name   string
	room   string
	writer *bufio.Writer
	reader *bufio.Reader
}

func (c *client) writeToConn(msg message) error {
	_, err := c.writer.WriteString(msg.String())
	if err != nil {
		return err
	}
	return c.writer.Flush()
}

func (c *client) readFromConn(send chan message) {
	for {
		data, err := c.reader.ReadString('\n')
		if err != nil {
			break
		}

		msg := message{
			from: c,
			name: c.name,
			room: c.room,
			text: strings.TrimSpace(data),
		}
		send <- msg
	}
}

func newClient(conn net.Conn) *client {
	c := &client{
		name:   makeName(),
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}

	return c
}

func makeName() string {
	fn := []string{
		"alpha", "bravo", "charlie", "delta", "echo", "foxtrot", "golf",
		"hotel", "india", "juliet", "kilo", "lima", "mike", "november",
		"oscar", "papa", "quebec", "romeo", "sierra", "tango", "uniform",
		"victor", "whiskey", "xray", "yankee", "zulu",
	}
	seed := time.Now().UnixNano()
	rn := rand.New(rand.NewSource(seed))

	p1 := rn.Intn(len(fn) - 1)
	p2 := rn.Intn(len(fn) - 1)
	return fn[p1] + "_" + fn[p2]
}
