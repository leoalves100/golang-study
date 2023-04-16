package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI            string
	Method         string
	Func           func(http.ResponseWriter, *http.Request)
	Authentication bool
}

// Configurar Adicionar todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	routes := routesUser
	routes = append(routes, routeLogin)

	for _, route := range routes {
		if route.Authentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authentication(route.Func)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}

	return r
}
