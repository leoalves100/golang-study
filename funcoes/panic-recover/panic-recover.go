// Ref: https://gobyexample.com/panic
package main

import "fmt"

func recuperar() {
	fmt.Println("Tentativa de recuperar a execução")

	// Recupera quando utilizado o panic()
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada!!!")
	}
}

func alunoStatusAprovado(n1, n2 int) bool {
	defer recuperar()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoStatusAprovado(6, 6))
	fmt.Println("PÓS EXECUÇÃO")
}
