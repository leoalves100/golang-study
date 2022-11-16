package main

import (
	"fmt"
)

func main() {

	// Verificação fora da função
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i: ", i)
	// 	time.Sleep(time.Second)
	// }

	// Verificação dentro da função
	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j: ", j)
	// 	time.Sleep(time.Second)
	// }

	// Usando range para interar
	nomes := [3]string{"leticia", "lucia", "leandri"}

	for i, nome := range nomes {
		fmt.Printf("Nome [%s] na posição [%v] \n", nome, i)
	}

	for i, nome := range "palavra" {
		fmt.Println(i, string(nome))
	}

	// Interando no map

	usuario := map[string]string{
		"nome": "leandro",
	}

	for i, nome := range usuario {
		fmt.Println(i, nome)
	}
}
