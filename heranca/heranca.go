package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura float32
	peso   uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("[Informações de pessoa]")

	var p1 = pessoa{"Leandro", 24, 1.80, 115}
	fmt.Println(p1)

	fmt.Println("[Informações de estudante]")
	var est1 = estudante{p1, "T.I", "FIAP"}
	fmt.Println(est1)
	fmt.Println(est1.nome)
}
