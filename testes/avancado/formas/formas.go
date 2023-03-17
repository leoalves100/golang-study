package formas

import (
	"fmt"
	"math"
)

// Cria um interface que implementa o método area
type Forma interface {
	Area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) Area() float64 {
	return (r.Altura) * (r.Largura)
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func main() {

	ret := Retangulo{10, 15}
	fmt.Println(ret)

	circ := Circulo{10}
	fmt.Println(circ)
}
