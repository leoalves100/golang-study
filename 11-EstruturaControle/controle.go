package main

import (
	"fmt"
)

func main() {

	numero := 10

	if numero > 15 {
		fmt.Println("Numero maior que 15", numero)
	} else {
		fmt.Println("Numero menor que 15", numero)
	}

	// Inicialização na condição
	if idade := 23; idade < 18 {
		fmt.Println("Você é menor de idade", idade)
	} else if idade > 18 {
		fmt.Println("Você é maior de idade", idade)
	} else {
		fmt.Println("Bugou!!!")
	}
}
