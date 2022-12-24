package server

import (
	"net/http"
	tp "web_server/types"
)

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

// esta funcion hace que el array de roules que el server tiene, cree un nuevo key y le asigne como valos una funcion
// recordar que rules es un map con llave y valor
func (s *Server) Handle(method, path string, handler http.HandlerFunc) {
	// verifica si existe el mapa de nivel de ruta, si no existe lo crea y luego crea la signacion del verbo http
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

// esta funcion permite agregar y manesjar middlewaress de forma dinamica
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...tp.Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
