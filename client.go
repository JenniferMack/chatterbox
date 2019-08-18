package chatterbox

import "bufio"

type client struct {
	write        *bufio.Writer
	roomMessages chan []byte
}

func (c *client) relay() {
	for byt := range c.roomMessages {
		_, err := c.write.Write(byt)
		if err != nil {
			return
		}
		c.write.Flush()
	}
}
