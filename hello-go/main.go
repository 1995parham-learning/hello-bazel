package main

import (
	"log"

	"github.com/go-fuego/fuego"

	"github.com/hello-bazel/hello-go/internal/handler"
)

func main() {
	s := fuego.NewServer()

	fuego.Get(s, "/", handler.HelloWorld)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
