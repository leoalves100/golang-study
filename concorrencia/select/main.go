// Ref: https://gobyexample.com/select
package main

import (
	"fmt"
	"time"
)

func main() {

	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {

		select {
		case msg1 := <-canal1:
			fmt.Println(msg1)
		case msg2 := <-canal2:
			fmt.Println(msg2)
		}

		/**
		*	Perdemos um tempo desnecessário por causa do canal2
		*	Ele leva 2seg para exibir o valor, enquanto o canal1 leva 0.5seg
		*	O canal1 poderia ser exibido 4x a mais que o canal2
		**/
		// Esperamos 0.5seg
		// msg1 := <-canal1
		// fmt.Println(msg1)

		// //Esperamos 2seg
		// msg2 := <-canal2
		// fmt.Println(msg2)
	}
}
