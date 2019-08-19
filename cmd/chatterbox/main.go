package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	chatterbox "github.com/JenniferMack/chatterbox/server"
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

	go chat.Collect("Lobby")
	<-quit
	println()
	log.Print("shutting down")
}
