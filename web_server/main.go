package main

import (
	hdlr "web_server/handlers"
	mw "web_server/middlewares"
	srvr "web_server/server"
)

func main() {
	server := srvr.NewServer(":3000")
	server.Handle("GET", "/", hdlr.HandleRoot)
	server.Handle("POST", "/home", hdlr.HandleHome)
	server.Handle("POST", "/user-create", hdlr.UserPostRequest)
	server.Handle("POST", "/generic-create", hdlr.PostRequest)
	server.Handle("POST", "/home", hdlr.HandleHome)
	server.Handle("POST", "/protected", server.AddMiddleware(hdlr.HandleHome, mw.CheckAuth(), mw.Logger()))
	server.Listen()
}
