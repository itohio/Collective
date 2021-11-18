package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	log.Println("Running")
	<-c
	log.Println("shutting down")
	os.Exit(0)
}
