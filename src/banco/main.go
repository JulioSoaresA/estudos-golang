package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoJulio := ContaCorrente{"Julio", 589, 123456, 125.50}

	fmt.Println(contaDoJulio)
}
