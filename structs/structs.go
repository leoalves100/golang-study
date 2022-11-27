/* Ref: https://youtu.be/Y4MKS3gJQ9Q
 *  Ref: https://gobyexample.com/structs
 */
package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("===Structs 1-Forma===")

	var u usuario
	u.nome = "Leandro Alves"
	u.idade = 24
	fmt.Println(u)

	fmt.Println("===Structs 2-Forma===")

	enderecoCompleto := endereco{"Rua Gandarela", 06}

	usuario2 := usuario{"Leandro", 24, enderecoCompleto}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Leandro Magalhães"}
	fmt.Println(usuario3)

	// Structs anônimo

	structAnonimo := struct {
		nome  string
		idade uint8
	}{
		nome:  "Leandro Alves",
		idade: 23,
	}

	fmt.Print("Struct Anônimo: \n", structAnonimo)
}
