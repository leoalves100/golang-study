// Ref: https://gobyexample.com/json
// Converter Struct ou Map em JSON

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	/** `json:"<key>"`
	/ Após ser convertido em json, <key> representa o nome da propriedade
	/**/
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Viralata", 3}
	fmt.Printf("Struct: %v\n", c)

	cachorroJSONConvert, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("JSON não convertido: %v", err)
	}

	// Retorna um slice de bytes
	fmt.Println("Slice de Bytes", cachorroJSONConvert)

	// Necessário utilizar função NewBuffer para ler o Array de Bytes
	res := bytes.NewBuffer(cachorroJSONConvert)
	fmt.Println("JSON: ", res)

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorroJSONConvertMap, err := json.Marshal(c2)
	if err != nil {
		log.Fatalf("JSON não convertido: %v", err)
	}

	// Necessário utilizar função NewBuffer para ler o Array de Bytes
	res2 := bytes.NewBuffer(cachorroJSONConvertMap)
	fmt.Println("JSON Map: ", res2)
}
