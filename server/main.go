package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/conley21p/AccelDatabase/server/internal/config"
	"github.com/conley21p/AccelDatabase/server/internal/server"
)

func main() {
	cfg := config.Load()
	s := server.NewServer(cfg)
	go func() {
		if err := s.Start(); err != nil {
			log.Fatalf("Server failed to start: %v\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Stop(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server shutdown complete")
}
