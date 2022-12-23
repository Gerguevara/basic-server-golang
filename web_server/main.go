package main

import server "web_server/server"

func main() {
	server := server.NewServer(":3000")
	server.Listen()
}
