package main

import "fmt"

func soma(num ...int) int {
	total := 0

	for _, value := range num {
		total += value
	}

	return total
}

func escrever(texto string, numero ...int) {

	for _, v := range numero {
		fmt.Println(texto, v)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5)
	fmt.Println(totalSoma)

	escrever("Olá mundo", 1, 2, 3, 4)
}
