package main

import (
	"fmt"
	"restapi/http"
	"restapi/src"
)

func main() {
	todoList := src.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start HTTP Server: ", err)
	}
}
