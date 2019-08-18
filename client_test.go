package chatterbox

import (
	"bufio"
	"net"
	"testing"
)

func TestWrite(t *testing.T) {
	in, out := net.Pipe()
	defer in.Close()
	defer out.Close()

	cl := client{
		write: bufio.NewWriter(out),
		recv:  make(chan []byte),
	}

	go cl.relay()
	defer close(cl.recv)

	want := "test message\n"
	cl.recv <- []byte(want)
	recv := bufio.NewReader(in)
	got, _ := recv.ReadString('\n')

	if got != want {
		t.Fatal(got)
	}
}

func TestRead(t *testing.T) {
	in, out := net.Pipe()
	defer in.Close()
	defer out.Close()

	cl := client{
		read: bufio.NewReader(out),
		send: make(chan []byte),
	}
	go cl.monitor()

	msg := "test\n"
	in.Write([]byte(msg))

	want := "> " + msg
	got := string(<-cl.send)
	if got != want {
		t.Fatal(got)
	}
}
