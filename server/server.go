package chatterbox

import "net"

func NewServer(port string) (net.Listener, error) {
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}

	return server, nil
}
