package chatterbox

import (
	"bufio"
	"net"
	"strings"
	"testing"
)

func TestWrite(t *testing.T) {
	in, out := net.Pipe()
	defer in.Close()
	defer out.Close()

	cl := client{
		writer: bufio.NewWriter(out),
	}

	go cl.writeToConn(message{name: "tester"})

	recv := bufio.NewReader(in)
	got, _ := recv.ReadString('\n')

	if !strings.Contains(got, "|tester>") {
		t.Fatal(got)
	}
}

func TestRead(t *testing.T) {
	in, out := net.Pipe()
	defer in.Close()
	defer out.Close()

	cl := client{
		name:   makeName(),
		reader: bufio.NewReader(out),
	}
	ch := make(chan message)
	go cl.readFromConn(ch)

	msg := "test\n"
	in.Write([]byte(msg))

	got := <-ch
	if !strings.Contains(got.String(), "> test") {
		t.Fatal(got.String())
	}
}
