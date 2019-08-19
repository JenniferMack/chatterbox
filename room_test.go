package chatterbox

import (
	"bufio"
	"net"
	"strings"
	"testing"
)

func TestNewRoom(t *testing.T) {
	r := NewRoom("lobby")
	_, conn := net.Pipe()

	r.connect(conn, 1)
	if l := len(r.clients); l != 1 {
		t.Fatal(l)
	}
}

func TestRoomMsg(t *testing.T) {
	r := NewRoom("lobby")
	in, out := net.Pipe()
	r.connect(in, 1)

	t.Run("room msg", func(t *testing.T) {
		want := "test message"
		go func() { r.messages <- []byte(want) }()
		got := <-r.clients[in].recv

		if string(got) != want {
			t.Fatalf("got: %s; want: %s", got, want)
		}
	})

	t.Run("output msg", func(t *testing.T) {
		want := "output test\r"
		go r.send([]byte(want))

		buf := bufio.NewReader(out)
		got, err := buf.ReadString('\r')
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(got, "output test") {
			t.Fatalf("got: %s; want: %s", got, want)
		}
	})
}
