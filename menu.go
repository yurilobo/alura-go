package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	// "io/ioutil"
	"bufio"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 50

func main() {

	exibeIntroducao()

	registraLog("site-falso", false)

	lerSitesDoArquivo()
	for {

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
	fmt.Println("Monitorando...")

	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}
func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)

	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}
func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "=online" + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
