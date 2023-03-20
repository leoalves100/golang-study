// Ref: https://gobyexample.com/channel-buffering

package main

import "fmt"

func main() {

	/**
	*	Adicionamos uma capacidade ao canal(buffer)
	*	Ao adicionar evitamos o bloqueio na transmissão da menssagem
	*	É bloqueado apenas quando atinge sua capacidade máxima(2)
	**/
	canal := make(chan string, 2)

	canal <- "Olá Mundo"

	mensagem := <-canal
	fmt.Println(mensagem)
}
