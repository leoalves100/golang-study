package main

import (
	"fmt"
	"introducao/endereco"
)

func main() {
	fmt.Println("===[Testando Tipos de Endereços]===")

	tipoEnderecoValido := endereco.TipoDeEndereco("Rua Paulista")
	fmt.Println("Tipo válido: ", tipoEnderecoValido)

	tipoEnderecoInvalido := endereco.TipoDeEndereco("Ruaa Paulista")
	fmt.Println("Tipo inválido: ", tipoEnderecoInvalido)
}
