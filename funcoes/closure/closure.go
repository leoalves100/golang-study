// Ref: http://goporexemplo.golangbr.org/closures.html#:~:text=Go%20suporta%20fun%C3%A7%C3%B5es%20an%C3%B4nimas%2C%20que,sem%20ter%20que%20nome%C3%A1%2Dlo.&text=Essa%20fun%C3%A7%C3%A3o%20intSeq%20retorna%20outra,i%20para%20formar%20o%20closure.
package main

import "fmt"

// Funções é um type em golang
func closure() func() string {
	txt := "Entrei na função closure()"

	funcao := func() string {
		return txt
	}

	return funcao
}

func main() {
	txt := "Dentro da função main()"
	fmt.Println(txt)

	funcaoNova := closure()
	fmt.Println(funcaoNova())
}
