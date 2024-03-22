package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// RouterConfig return routes configurations
func RouterConfig() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
