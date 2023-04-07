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

	var comando int

	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi: ", comando)

}
