package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const VERSAO = 1.1
const monitoramentos = 3
const delay = 5

func main() {
	fmt.Println("Este programa está na versão", VERSAO)
	fmt.Println()
	for {
		exibeIntroducao()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Fechando o programa...")
			os.Exit(0)
		default:
			os.Exit(-1)
			fmt.Println("Não conheço este comando")
		}
	}
}

func exibeIntroducao() {
	fmt.Println("Escolha uma das opções abaixo:")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("")
}

func leComando() int {
	fmt.Print("Resposta: ")
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println()
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando Monitoramento...")
	// var sites [4]string
	// sites[0] = "https://www.alura.com.br"...

	sites := []string{
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
	}

	// índice, valor (retorno "range")
	for i := 0; i < monitoramentos; i++ {
		for _, v := range sites {
			testaSite(v)
		}
		if i != monitoramentos-1 {
			time.Sleep(delay * time.Second)
			fmt.Println("")
		}
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)
	switch resp.StatusCode {
	case 200:
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	default:
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
	}
}
