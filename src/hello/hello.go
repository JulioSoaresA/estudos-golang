package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string
	versao := 1.1
	fmt.Println("Digite seu nome:")
	fmt.Scan(&nome)

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão ", versao)

	fmt.Println("O tipo da variavel nome é:", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel versao é:", reflect.TypeOf(versao))

	fmt.Println("\n1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("Escolha uma opção: ")

	var comando int

	fmt.Scan(&comando)

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
	default:
		fmt.Println("Comando inválido.")
	}

}
