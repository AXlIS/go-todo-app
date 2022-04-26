package main

import (
	"github.com/AXlIS/go-todo-app"
	"github.com/AXlIS/go-todo-app/pkg/handler"
	"log"
)

func main() {
	server := new(todo.Server)
	handlers := new(handler.Handler)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
