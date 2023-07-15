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

// Configurar Add all routes inside the router
func Configurar(r *mux.Router) *mux.Router {
	routes := routesUser
	routes = append(routes, routeLogin)
	routes = append(routes, routesPublications...)

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
