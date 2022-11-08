package main

import "fmt"

// Função de soma, 1-Forma
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Função com 2 retornos
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// Função de print, 2-Forma
	var lambda = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	lambda("Teste de função lambda")

	// Ignorar retorno _
	resultSoma, _ := calculosMatematicos(10, 5)
	fmt.Println(resultSoma)
}
