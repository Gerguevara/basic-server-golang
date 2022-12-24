package server

import (
	"net/http"
)

// solo es  una simple estructura que contiene reglas de nueestro routing
type Router struct {
	// rules map[string]http.HandlerFunc
	//ahora tenemos un mapa de mapas para manejar los verbos HTTP, un mapa que apunta otro mapa y este a los handlers
	rules map[string]map[string]http.HandlerFunc
}

// el & y el * nos permite evitar crear copias,
func NewRouter() *Router {
	return &Router{
		//acca creamos un map vacio dopnde iremos poniendo las rutas
		//make puede crear en una sola sentencia el map de mapas
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// este metodo se crea para cumplir cona las reglas de la interfaz handlrer que es necesario que se cumpla
// para poder usar el router en metodo Listen y http.Handle
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// de esta forma encontramos segun la ruta la funcion que debera pasar
	handler, methodExist, exist := r.FindHandler(request.URL.Path, request.Method)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}

// busca el handler segun el path, retorna 2 valores la funcion mapeada a esa ruta y si la funcion exite
func (r *Router) FindHandler(path, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}
