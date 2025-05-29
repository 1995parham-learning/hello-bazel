package handler

import (
	"fmt"

	"github.com/go-fuego/fuego"
)

func HelloWorld(c fuego.ContextNoBody) (string, error) {
	to := c.QueryParam("to")
	if to == "" {
		to = "World"
	}

	return fmt.Sprintf("Hello, %s!", to), nil
}
