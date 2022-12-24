package types

import "net/http"

// solo es la interface de un middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc

//un middleware es basicamente un funcion que recibe una funcion, ejecuta su propia logica
// luego retorna la funcion que recibio,
