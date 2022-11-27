package main

import "fmt"

type usuario struct {
	Nome  string
	Idade uint8
}

// Metodo salvar de usuario
func (u usuario) salvar() {
	fmt.Printf("Salvado dados do usuario %s...\n", u.Nome)
}

// Utilizando ponteiros nos métodos
func (u *usuario) fazerNiver() {
	u.Idade++
}

func main() {
	usuario := usuario{
		Nome:  "Leandro",
		Idade: 23,
	}

	fmt.Println(usuario)

	usuario.salvar()

	usuario.fazerNiver()
	fmt.Println(usuario.Idade)
}
