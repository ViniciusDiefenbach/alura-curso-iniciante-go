package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const VERSAO = 1.1
const monitoramentos = 5
const delay = 5

func main() {
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
	fmt.Println("Este programa está na versão", VERSAO)
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando Monitoramento...")
	// var sites [4]string
	// sites[0] = "https://www.alura.com.br"...

	sites := []string{"https://www.alura.com.br", "https://www.caelum.com.br"}

	// índice, valor (retorno "range")
	for i := 0; i < monitoramentos; i++ {
		for _, v := range sites {
			testaSite(v)
		}
		if i != monitoramentos-1 {
			time.Sleep(delay * time.Second)
		}
	}
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
