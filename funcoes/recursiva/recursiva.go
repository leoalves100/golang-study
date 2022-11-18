// Ref: https://gobyexample.com/recursion
package main

import "fmt"

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	position := uint(10)

	res := fibonacci(position)

	fmt.Printf("Posição do numero %v é %v\n", position, res)
}
