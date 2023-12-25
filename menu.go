package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Yuri"
	versao := 1.1

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Estamos na versão: ", versao)
	fmt.Println("Estamos usando o tipo : ", reflect.TypeOf(versao), " em versao")
	fmt.Println("1-Iniciar Monitoramento ")
	fmt.Println("2-Exibir os logs ")
	fmt.Println("0-Sair do programa ")

	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O endeco da minha variavel comando é", &comando)
	fmt.Println("O comando escolhido foi ", comando)
}
