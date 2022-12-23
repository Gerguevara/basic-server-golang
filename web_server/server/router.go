package server

import (
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
	// de esta forma encontramos segun la ruta la funcion que debera pasar
	handler, exist := r.FindHandler(request.URL.Path)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler(w, request)
}

// busca el handler segun el path, retorna 2 valores la funcion mapeada a esa ruta y si la funcion exite
func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
}
