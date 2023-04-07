package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const monitoramentos = 2
const delay = 2

func main() {
	exibeIntro()

	for {
		exibeMenu()

		comando := lerComando()

		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo Logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do Programa...")
		// } else {
		// 	fmt.Println("Comando inválido.")
		// }

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido.")
			os.Exit(-1)
		}
	}
}

func exibeIntro() {
	var nome string
	versao := 1.1
	fmt.Println("Digite seu nome:")
	fmt.Scan(&nome)

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão ", versao)

	fmt.Println("O tipo da variavel nome é:", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel versao é:", reflect.TypeOf(versao))
}

func exibeMenu() {
	fmt.Println("")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("Escolha uma opção: ")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func iniciarMonitoramento() {
	sites := []string{
		"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
		"https://www.google.com.br",
	}

	fmt.Println("Monitorando...")

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso.")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
