package main

import (
	"fmt"
	"os"
)

var VERSAO = 1.1

func main() {
	exibeIntroducao()
	comando := leComando()
	switch comando {
	case 1:
		fmt.Println("Iniciando Monitoramento...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 3:
		fmt.Println("Fechando o programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
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
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}
