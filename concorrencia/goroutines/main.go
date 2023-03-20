// Ref: https://gobyexample.com/goroutines
// Ref: https://medium.com/trainingcenter/goroutines-e-go-channels-f019784d6855
package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!")
	escrever("Programando em Golang!")
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
