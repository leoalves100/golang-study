package main

import "fmt"

func main() {

	//Explícito
	var nome string = "Leandro Alves"
	fmt.Println(nome)

	//Implícito
	idade := 23 + 1
	fmt.Println(idade)

	//Declaração conjunta

	// 1 - Modelo
	var (
		sexo          string = "Masculino"
		planetaOrigem string = "Marte"
	)
	fmt.Println(sexo, planetaOrigem)

	// 2 - Modelo
	profissao, cargo := "SRE", "Pleno"
	fmt.Println(profissao, cargo)

	// Constantine
	const constantine string = "Maluco brabo!!!"
	fmt.Println(constantine)

	// Inverter valores
	profissao, cargo = cargo, profissao
	fmt.Println(profissao, cargo)
}
