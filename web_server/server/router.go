package server

import (
	"fmt"
	"net/http"
)

// solo es  una simple estructura que contiene reglas de nueestro routing
type Router struct {
	rules map[string]http.HandlerFunc
}

// el & y el * nos permite evitar crear copias,
func NewRouter() *Router {
	return &Router{
		//acca creamos un map vacio dopnde iremos poniendo las rutas
		rules: make(map[string]http.HandlerFunc),
	}
}

// este metodo se crea para cumplir cona las reglas de la interfaz handlrer que es necesario que se cumpla
// para poder usar el router en metodo Listen y http.Handle
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	//impresion de mensaje respuesta que el servidor da a la ruta
	//Fprintf es un escritor, que recive w que es el escritor asignado, y el mensaje que queremos mostrar
	fmt.Fprintf(w, "it's working")
}
