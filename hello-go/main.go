package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-fuego/fuego"

	"github.com/hello-bazel/hello-go/internal/handler"
)

func main() {
	s := fuego.NewServer()

	fuego.Get(s, "/", handler.HelloWorld)

	// Create context that listens for interrupt signals
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Run server in goroutine
	go func() {
		if err := s.Run(); err != nil {
			log.Printf("server error: %v", err)
		}
	}()

	log.Println("server started, press Ctrl+C to stop")

	// Wait for interrupt signal
	<-ctx.Done()
	log.Println("shutting down gracefully...")
}
