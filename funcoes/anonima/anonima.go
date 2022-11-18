package main

import "fmt"

func main() {

	// Declara a função e executa logo em seguida
	retorno := func(txt string) string {
		return fmt.Sprintf("Recebido -> %s", txt)
		// Adiciona parâmetro
	}("Passando paramêtro")

	fmt.Printf(retorno)
}
