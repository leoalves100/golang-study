package main

import (
	"crud/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const PORT string = ":5000"

	router := mux.NewRouter()

	router.HandleFunc("/usuarios", repository.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", repository.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuario/{id}", repository.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuario/{id}", repository.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuario/{id}", repository.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Printf("Escutando na porta %v", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
