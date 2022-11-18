// Ref: https://gobyexample.com/pointers
package main

import "fmt"

func invertSinal(num int) int {
	return num * -1
}

func invertSinalPonteiro(num *int) int {
	*num = *num * -1
	return *num
}

func main() {
	num := 20
	numInvert := invertSinal(num)
	fmt.Println("Cópia da variável: ", numInvert)
	fmt.Println("Cópia da variável: ", num)

	numNew := 40
	numInvertPonteiro := invertSinalPonteiro(&numNew)
	fmt.Println("Ponteiro: ", numInvertPonteiro)
	fmt.Println("Ponteiro: ", numNew)
}
