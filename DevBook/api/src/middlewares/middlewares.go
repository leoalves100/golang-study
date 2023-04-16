// # Ref: https://drstearns.github.io/tutorials/gomiddleware/
package middlewares

import (
	"api/src/authentication"
	"api/src/response"
	"log"
	"net/http"
)

// Logger Print information requests
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s\n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authentication check user request in authentication
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidToken(r); err != nil {
			response.Erro(w, http.StatusUnauthorized, err)
		}
		next(w, r)
	}
}
