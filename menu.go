package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	exibeMenu()
	// _, idade := devolveNomeEIdade()
	// fmt.Println("e tenho", idade, "anos")

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
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheco este comando")
		os.Exit(-1)

	}

}

// func devolveNomeEIdade() (string, int) {
// 	nome := "Lobo"
// 	idade := 28
// 	return nome, idade
// }

func exibeIntroducao() {
	nome := "Yuri"
	versao := 1.1

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Estamos na versão: ", versao)
	fmt.Println("Estamos usando o tipo : ", versao, " em versao")

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

func iniciarMonitoramento() {
	fmt.Println("Monitorando..")
	site := "https://www.alura.com.br/"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}
