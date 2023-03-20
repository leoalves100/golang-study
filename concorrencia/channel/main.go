// Ref: https://gobyexample.com/channels

package main

import (
	"fmt"
	"time"
)

func main() {

	// Cria um canal
	canal := make(chan string)

	go escrever("Tô na área", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	/** Recebimento da mensagem no canal é uma operação bloqueante
	* 	O go resperar o recebimento da mensagem para continuação da execução
	**/

	// 1 - Exibe o valor do canal
	for {
		/** fatal error: all goroutines are asleep - deadlock!
		*   Acontece quando não existe mais mensagens sendo trafegadas no canal
		*	Porém ele continua aberto, esperando receber novos dados
		**/

		mensagem, statusCanal := <-canal
		/**
		* Verificar se o canal está aberto, caso contrário, saí do loop
		**/
		if !statusCanal {
			break
		}
		fmt.Println(mensagem)
	}

	// 2 - Exibe o valor do canal (Forma secundária)
	for menssagem := range canal {
		fmt.Println(menssagem)
	}

	fmt.Println("Fim do programa!!!")
}

func escrever(texto string, canal chan string) {
	for i := 1; i <= 5; i++ {
		// Envia o valor para dentro do canal
		canal <- fmt.Sprintf("Texto: %v [%v]", texto, i)
		time.Sleep(time.Second)
	}

	/**
	*	Fecha o canal para enviar e receber dados
	**/
	close(canal)
}
