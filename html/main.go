package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type usuarios struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!!!"))
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := usuarios{
			"leandro",
			"leandro.alves351@gmail.com",
		}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Listen: http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
