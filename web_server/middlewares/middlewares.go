package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// primer middleware
func CheckAuth() Middleware {
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
func Logger() Middleware {
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
