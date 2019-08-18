package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/JenniferMack/chatterbox"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	chat, err := chatterbox.NewServer("5050")
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
