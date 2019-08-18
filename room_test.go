package chatterbox

import (
	"bufio"
	"net"
	"testing"
)

func TestNewRoom(t *testing.T) {
	r := NewRoom("lobby")
	_, conn := net.Pipe()

	r.Connect(conn)
	if l := len(r.clients); l != 1 {
		t.Fatal(l)
	}
}

func TestRoomMsg(t *testing.T) {
	r := NewRoom("lobby")
	in, out := net.Pipe()
	r.Connect(in)

	t.Run("room msg", func(t *testing.T) {
		want := "test message"
		go func() { r.messages <- []byte(want) }()
		got := <-r.clients[in].roomMessages

		if string(got) != want {
			t.Fatalf("got: %s; want: %s", got, want)
		}
	})

	t.Run("output msg", func(t *testing.T) {
		want := "output test\r"
		go r.Send([]byte(want))

		buf := bufio.NewReader(out)
		got, err := buf.ReadString('\r')
		if err != nil {
			t.Fatal(err)
		}

		if got != want {
			t.Fatalf("got: %s; want: %s", got, want)
		}
	})
}
