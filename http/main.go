package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!!!"))
}

func main() {

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
