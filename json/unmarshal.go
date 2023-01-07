// Converter JSON em Struct

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type cachorro struct {
	/** `json:"<key>"`
	/ Após ser convertido em json, <key> representa o nome da propriedade
	/ Para ignorar uma variável no json, passar assim: `json:"-"`
	/**/
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	// Convert JSON em Struct
	cachorroEmJSON := `{"nome":"Rex","raca":"Viralata","idade":3}`

	var c cachorro

	// Unmarshal recebe um slice de bytes
	// Necessário converter
	if err := json.Unmarshal([]byte(cachorroEmJSON), &c); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tipo da variavel c: %v\n", reflect.TypeOf(c))
	fmt.Println("Struct: ", c)

	// Convert JSON em Map
	cachorroEmJSON2 := `{"nome":"Tody","raca":"não definido","idade":5}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroEmJSON2), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Printf("Tipo da variavel c2: %\n", reflect.TypeOf(c2))
	fmt.Println("Struct: ", c2)
}
