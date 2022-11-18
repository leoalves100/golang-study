package main

import "fmt"

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "numero é 1"
		// Pula para o outro case
		fallthrough
	case numero == 2:
		diaDaSemana = "numero é 2"
	default:
		return "numero inválido\n"
	}

	return diaDaSemana
}

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo\n"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	default:
		return "Número inválido"
	}
}

func main() {
	fmt.Println("===[Switch]===")

	fmt.Printf(diaDaSemana(1))

	fmt.Printf(diaDaSemana2(2))
}
