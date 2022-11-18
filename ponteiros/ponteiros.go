package main

import "fmt"

func main() {
	fmt.Println("[Ponteiros]")

	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	// Ponteiro - Referência de memória
	var variavel3 int = 100
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, *ponteiro) // Deresferenciação

	variavel3 = 200
	fmt.Println(variavel3, *ponteiro)
}
