package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/JenniferMack/chatterbox"
)

func main() {
	var flagPort = flag.String("port", "5050", "server port")
	flag.Parse()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	chat, err := chatterbox.NewServer(*flagPort)
	if err != nil {
		log.Fatal(err)
	}
	defer chat.Close()

	room := chatterbox.NewRoom("Lobby")
	defer room.Close()
	go room.Accept(chat)

	<-quit
	println()
	log.Print("shutting down")
}
