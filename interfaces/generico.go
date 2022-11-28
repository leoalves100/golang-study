// Ref: https://blog.streamelements.com/an-introduction-to-generics-in-go-cc8cdae15ef2

package main

import "fmt"

type generico interface{}

func generic(g generico) {
	fmt.Printf("Seu valor é: %v %t \n", g, g)
}

func main() {
	generic("Leandro")
	generic(24)
	generic(24.2)

	// Map genérico
	mapa := map[interface{}]interface{}{
		1:        "string",
		1.1:      true,
		"string": 100,
	}

	fmt.Println(mapa)
}
