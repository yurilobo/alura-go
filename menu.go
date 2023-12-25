package main

import (
	"fmt"
	"reflect"
)

func main() {

	exibeIntroducao()

	exibeMenu()

	comando := lerComando()

	// if comando == 1 {
	// 	fmt.Println("Monitorando..")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo logs")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa")
	// } else {
	// 	fmt.Println("Não conheco este comando")
	// }
	switch comando {
	case 1:
		fmt.Println("Monitorando..")
	case 2:
		fmt.Println("Exibindo logs")
	case 0:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Não conheco este comando")
	}

}
func exibeIntroducao() {
	nome := "Yuri"
	versao := 1.1

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Estamos na versão: ", versao)
	fmt.Println("Estamos usando o tipo : ", reflect.TypeOf(versao), " em versao")

}
func exibeMenu() {
	fmt.Println("1-Iniciar Monitoramento ")
	fmt.Println("2-Exibir os logs ")
	fmt.Println("0-Sair do programa ")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi ", comandoLido)
	return comandoLido
}
