package chatterbox

import "testing"

func TestNewServer(t *testing.T) {
	srv, err := NewServer("5050")
	if err != nil {
		t.Fatal(err)
	}
	if err := srv.Close(); err != nil {
		t.Fatal(err)
	}
}
