package middlewares

import "net/http"

// solo es la interface de un middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc
