// Ref: https://gobyexample.com/waitgroups
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	// Quantidade de goroutines que será executadas
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo!")
		// Remove um goroutine do Add(2)
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Golang!")
		waitGroup.Done()
	}()

	// Espera as goroutines serem finalizadas
	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
