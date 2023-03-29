package rotas

import (
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
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
