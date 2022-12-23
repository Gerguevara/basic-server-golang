package server

import "net/http"

type Server struct {
	port string
	//apuntor al archivo de router.go
	router *Router
}

// retorna un server como tal, la referencia de un server, (es el contructor)
// lo que retorna es el server cono tal no copias
func NewServer(port string) *Server {
	return &Server{
		port: port,
		//aqui es donde el guarda las rutas, (inicia como un array vacio)
		router: NewRouter(),
	}
}

// reciber function para habilitar escuchar peticiones

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
