// Ref: https://gobyexample.com/worker-pools
package main

import (
	"fmt"
)

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

/**
*	Podemos especificar a função do canal, recebimento ou envio de menssagens
*	tarefas <-chan int = só recebe dados
*	resultados chan<- int = só envia dados
**/
func worker(tarefas <-chan int, resultados chan int) {

	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}

}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}
