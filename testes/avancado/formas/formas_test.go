package avancado

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}

		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {

			// Fatal: Para a execução dos testes ao identificar erro
			t.Fatalf("A área recebida %f é diferente da esperada: %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {

			// Fatal: Para a execução dos testes ao identificar erro
			t.Fatalf("A área recebida %f é diferente da esperada: %f", areaRecebida, areaEsperada)
		}
	})
}
