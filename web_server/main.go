package main

import (
	handlers "web_server/handlers"
	server "web_server/server"
)

func main() {
	server := server.NewServer(":3000")
	server.Handle("/", handlers.HandleRoot)
	server.Handle("/home", handlers.HandleHome)
	server.Listen()
}
