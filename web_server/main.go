package main

import (
	hdlr "web_server/handlers"
	mw "web_server/middlewares"
	s "web_server/server"
)

func main() {
	server := s.NewServer(":3000")
	server.Handle("/", hdlr.HandleRoot)
	server.Handle("/home", hdlr.HandleHome)
	server.Handle("/protected", server.AddMiddleware(hdlr.HandleHome, mw.CheckAuth(), mw.Loggin()))
	server.Listen()
}
