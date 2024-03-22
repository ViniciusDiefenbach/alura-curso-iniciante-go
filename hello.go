package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			imprimeLogs()
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
	fmt.Println()
	// var sites [4]string
	// sites[0] = "https://www.alura.com.br"...

	sites := leSitesDoArquivo()

	// índice, valor (retorno "range")
	for i := 0; i < monitoramentos; i++ {
		for _, v := range sites {
			testaSite(v)
		}
		if i != monitoramentos-1 {
			fmt.Println("")
			time.Sleep(delay * time.Second)
		}
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	switch resp.StatusCode {
	case 200:
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	default:
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	// arquivo, err := os.ReadFile("sites.txt")
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		sites = append(sites, strings.TrimSpace(linha))
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	fmt.Println("Exibindo Logs...")
	fmt.Println()

	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
