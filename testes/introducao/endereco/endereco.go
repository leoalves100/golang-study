package endereco

import (
	"strings"
)

// TipoDeEndereco Verifica se um endereço é valido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoMinuscula, " ")[0]

	enderecoTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTipoValido = true
		}
	}

	if enderecoTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return endereco
}
