package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else if valorDoSaque < 0 {
		return "Valor inválido"
	} else {
		return "saldo insuficiente"
	}
}

func main() {
	contaDoJulio := ContaCorrente{}
	contaDoJulio.titular = "Julio"
	contaDoJulio.numeroAgencia = 589
	contaDoJulio.numeroConta = 123456
	contaDoJulio.saldo = 125.50

	fmt.Println(contaDoJulio.saldo)

	contaDaMonalisa := ContaCorrente{}
	contaDaMonalisa.titular = "Monalisa"
	contaDaMonalisa.numeroAgencia = 589
	contaDaMonalisa.numeroConta = 123457
	contaDaMonalisa.saldo = 500

	fmt.Println(contaDoJulio.Sacar(-100))
	fmt.Println(contaDoJulio.saldo)

}
