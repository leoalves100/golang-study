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
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Func).Methods(rota.Method)
	}

	return r
}
