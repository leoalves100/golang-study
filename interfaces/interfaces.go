/*
 * Ref: https://gobyexample.com/interfaces
 * Ref: https://www.youtube.com/watch?v=2zTENBJTlD0
 */
package main

import (
	"fmt"
	"math"
)

// Cria um interface que implementa o método area
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("Área da forma é: %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return (r.altura) * (r.largura)
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {

	retangulo := retangulo{10, 15}
	escreverArea(retangulo)

	circulo := circulo{10}
	escreverArea(circulo)
}
