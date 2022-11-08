package main

import (
	"fmt"
)

func main() {

	fmt.Println("Array")

	// 1-Forma de declarar e preencher array
	var array1 [5]string
	array1[0] = "Possição 0"
	fmt.Println("1-Forma: ", array1)

	// 2-Forma de declarar e preencher array
	array2 := [2]string{"Pos 1", "Pos 2"}
	fmt.Println("2-Forma: ", array2)

	// Seta tamanho do array de forma dinâmica
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array com tamanho dinâmico: ", array3)

	// Cut array
	fmt.Println("Cut array:", array3[0:2])

	array3[0] = 10
	fmt.Println("Posição alterada:", array3[0:2])

	fmt.Println("Slice")

	// 1-Froma de declarar e preencher um slice
	slice1 := []int{10, 20, 30}
	fmt.Println("1-Forma:", slice1)

	// 2-Froma de declarar e preencher um slice
	slice2 := []int{}
	slice2 = append(slice2, 40)
	fmt.Println("2-Forma:", slice2)

	// Array internos
	slice3 := make([]float32, 10, 11)
	fmt.Println("Usando make: ", slice3)
	fmt.Println("Tamanho: ", len(slice3))
	fmt.Println("Capacidade: ", cap(slice3))

	slice3 = append(slice3, 10)
	slice3 = append(slice3, 11)
	fmt.Println("Ultrapasando a capacidade: ", slice3)
	fmt.Println("Tamanho: ", len(slice3))
	fmt.Println("Capacidade: ", cap(slice3))
}
