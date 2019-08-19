# chatterbox
Simple full-duplex TCP chat server

## Building
Via make:

`make`

Via Go tools:

`go build ./cmd/chatterbox/...`

## Running
From the root directory:
`./chatterbox`

Changing the default port:
`./chatterbox -port 8080`

## Usage
Connecting with telnet:
`telnet localhost 5050`

Closing the chat:
`/close`

Changing the room:
`/room <name>`

## Docker
The makefile can build, run and stop a chatterbox docker container.

To build:
`make docker`

To run:
`make docker-run`

To stop:
`make docker-kill`

