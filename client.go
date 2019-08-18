package chatterbox

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

type client struct {
	name  string
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
		data = strings.TrimSpace(data)
		msg := fmt.Sprintf("<%s|%s> %s\n",
			time.Now().UTC().Format("15:04"), c.name, data)
		c.send <- []byte(msg)
	}
}
