package main

import (
	"fmt"
)

func main() {

	fmt.Println("===[Maps]===")

	usuario := map[string]string{
		"nome":      "Leandro",
		"sobrenome": "Alves",
	}

	fmt.Println(usuario)

	// Acessando propriedade
	fmt.Println("Acessando propriedade: ", usuario["nome"])

	usuario2 := map[string]map[string]string{
		"carac": {
			"cabelo": "preto",
			"olhos":  "castanhos",
		},
		"habilidades": {
			"comida": "miojo",
			"aws":    "melhor que azure",
		},
	}

	fmt.Println("[map] aninhado: ", usuario2)
	delete(usuario2, "habilidades")
	fmt.Println("[map] aninhado: ", usuario2)

}
