package main

import "fmt"

// Comentário na função
func calculosMatemeticos(n1 int, n2 int) (soma int, subtracao int) {
	/* Não precisamos declarar o tipo,
	   pois foi adicionado no retorno
	*/
	soma = n1 + n2
	subtracao = n1 - n2

	return
}

func main() {
	soma, subtracao := calculosMatemeticos(10, 20)
	fmt.Printf("Soma: %v Subtração: %v \n", soma, subtracao)
}
