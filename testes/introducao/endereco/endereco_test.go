package endereco_test

import (
	. "introducao/endereco"
	"testing"
)

type cenarioDeTest struct {
	enderecoInserido string
	retornoEsperado  string
}

// Cenário de testes
func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()

	cenarioDeTestes := []cenarioDeTest{
		{"Avenida Paulista", "Avenida"},
		{"Rua da 500", "Rua"},
		{"Estrada da Morte, n* 20", "Estrada"},
		{"Rodovia BR-090", "Rodovia"},
		{"RUA DOS BOBOS", "Rua"},
		// {"Viela da Quinta", "Viela"},
		// {"Travessa dos Alfedeiros", "Travessa"},
	}

	for _, cenario := range cenarioDeTestes {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("Endereço inválido, Valor: %s", cenario.retornoEsperado)
		}
	}
}

func TestParapelo(t *testing.T) {
	t.Parallel()

	if 2 > 3 {
		t.Error("Teste quebrou!")
	}
}

/** Assinatura da função de teste
 *  Começa com nome Test
 *  Precisa ser uma função publica
**/

// Primeiro modelo de teste
// func TestTipoDeEndereco(t *testing.T) {
// 	enderecoParaTeste := "Avenidaa Paulista"

// 	tipoDeEnderecoEsperado := "Avenida"

// 	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

// 	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
// 		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s | Recebeu %s",
// 			tipoDeEnderecoEsperado,
// 			tipoDeEnderecoRecebido,
// 		)
// 	}
// }
