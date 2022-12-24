// los handlers nos siver par buscasr qeu funcion esta relacionada a que ruta
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	tp "web_server/types"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wellcome to Home")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	//la metadata es una interface generica usada como comodin que se usar par setear todos los campos ahi
	var metadata tp.Metadata
	// enviamos la referencia de esta variable
	err := decoder.Decode(&metadata)

	if err != nil {
		fmt.Fprintf(w, "error %v", err)
		return
	}
	//haciendo esto la interface da un error y no nos permite manipular propiedades de la interface generica
	// fmt.Println(metadata["name"])

	fmt.Fprintf(w, "Payload %v\n", metadata)

}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user tp.User

	err := decoder.Decode(&user)

	if err != nil {
		fmt.Fprintf(w, "error %v", err)
		return
	}

	response, err := user.ToJson()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//respondia  al cliente con un string
	// fmt.Fprintf(w, "Payload %v\n", user)
	w.Header().Set("Content-Type", "application/json")

	//el writer envia la respuesta al cliente
	w.Write(response)
}
