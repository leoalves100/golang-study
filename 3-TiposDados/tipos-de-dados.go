package main

import "fmt"

func main() {

	//Tipos de inteiros
	// - int Arch do PC
	// - int 8
	// - int 16
	// - int 32
	// - int 64

	// Apenas inteiros positivos
	// - uint
	// - uint8
	// - uint16
	// - uint32
	// - uint64
	// - uintptr

	// Alias
	// INT32 = RUNE
	var number rune = 123456789
	fmt.Printf("RUNE == INT32: %v\n", number)

	// BYTE = UINT8
	var numberUnit byte = 100
	fmt.Printf("BYTE == UINT8: %v\n", numberUnit)

	// Tipos de Float
	// -float32
	// -float64
	var numberFloat32 float32 = 100.01
	fmt.Printf("FLOAT32: %v\n", numberFloat32)

	var numberFloat64 float32 = 100.01123
	fmt.Printf("FLOAT64: %v\n", numberFloat64)

	var str string = "Texto"
	fmt.Printf("STRING: %v\n", str)

	char := 'T'
	fmt.Println(char)
}
