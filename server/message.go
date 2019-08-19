package chatterbox

import "time"

type message struct {
	text string
	room string
	name string
	from *client
}

func (m *message) String() string {
	t := time.Now().UTC().Format("15:04")
	return "<" + t + "|" + m.name + "> " + m.text + "\n"
}
