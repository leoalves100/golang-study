package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canal := multiplexar(
		escrever("Ficou acordado até tarde secando o mengão"),
		escrever("Fiquei piranha"),
	)

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

/**
*	Recebe dois canais com entrada e retorna um canal como saída
*	Dessa forma, conseguimos criar uma função multiplexado
*	Unificando todos os dados em um único canal
**/
func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
