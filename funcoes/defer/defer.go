// Ref: https://gobyexample.com/defer
package main

import "fmt"

func primeiraFuncao() {
	fmt.Println("Executando primeira função")
}

func segundaFuncao() {
	fmt.Println("Executando segunda função")
}

func calculaMedia(n1, n2 int) bool {
	fmt.Println("Entrando na função calculaMedia()")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {

	// Adia a execução da função
	defer primeiraFuncao()

	segundaFuncao()

	fmt.Println(calculaMedia(10, 6))
}
