package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("Worker started. Press Ctrl-C to stop server")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	active := true

	go func() {
		signalType := <-ch
		signal.Stop(ch)
		log.Println("Signal type : ", signalType)
		active = false
	}()

	for active {
		log.Println("doing a task...")
		time.Sleep(1000 * time.Millisecond)
	}

	log.Println("Exiting...")
}
