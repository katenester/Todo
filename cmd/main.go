package main

import (
	"github.com/katenester/Todo"
	"github.com/katenester/Todo/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo_app.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
