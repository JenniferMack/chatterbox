package chatterbox

import (
	"bufio"
)

type client struct {
	write *bufio.Writer
	read  *bufio.Reader
	send  chan []byte
	recv  chan []byte
}

func (c *client) relay() {
	for byt := range c.recv {
		_, err := c.write.Write(byt)
		if err != nil {
			return
		}
		c.write.Flush()
	}
}

func (c *client) monitor() {
	for {
		data, err := c.read.ReadString('\n')
		if err != nil {
			return
		}
		c.send <- []byte("> " + data)
	}
}
