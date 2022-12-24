package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"
	md "web_server/types"
)

// primer middleware
func CheckAuth() md.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking oauth credentials")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

// otro middleware
func Logger() md.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			//logica del middleware
			start := time.Now()
			// Defer una funcion anonima para que se ejecute justo al final de esta funcion
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			//llamamos la siquiente funcion
			f(w, r)
		}
	}
}
