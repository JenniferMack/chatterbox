package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	println()
	log.Print("shutting down")
}
