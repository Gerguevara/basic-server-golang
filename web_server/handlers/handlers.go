// los handlers nos siver par buscasr qeu funcion esta relacionada a que ruta
package handlers

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wellcome to Home")
}
