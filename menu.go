package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	exibeMenu()

	comando := lerComando()

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
	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "esta com problemas!", resp.StatusCode)
	}

}
