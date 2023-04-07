package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	exibeIntro()

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
		fmt.Println("Monitorando...")
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
	fmt.Println("\n1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("Escolha uma opção: ")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}
