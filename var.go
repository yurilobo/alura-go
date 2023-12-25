package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Yuri"
	var versao = 1.1
	var idade1 int
	var idade2 = 28
	fmt.Println("Olá Sr.", nome, "nome, sua idade é:", idade1, "ou", idade2)
	fmt.Println("Estamos na versão: ", versao)
	fmt.Println("Estamos usando o tipo : ", reflect.TypeOf(versao), " em versao")
}
