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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso. Novo saldo é de", c.saldo
	} else {
		return "Falha ao realizar depósito. Valor inválido", valorDoDeposito
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) string {
	if valorDaTransferencia > 0 && valorDaTransferencia <= c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferência realizada com sucesso"
	} else {
		return "Falha ao realizar transferência."
	}
}

func main() {
	contaDoJulio := ContaCorrente{}
	contaDoJulio.titular = "Julio"
	contaDoJulio.numeroAgencia = 589
	contaDoJulio.numeroConta = 123456
	contaDoJulio.saldo = 125.50

	fmt.Println("Saldo da conta do Julio: ", contaDoJulio.saldo)

	contaDaMonalisa := ContaCorrente{}
	contaDaMonalisa.titular = "Monalisa"
	contaDaMonalisa.numeroAgencia = 589
	contaDaMonalisa.numeroConta = 123457
	contaDaMonalisa.saldo = 500

	fmt.Println("Saldo da conta da Monalisa: ", contaDaMonalisa.saldo)

	status := contaDoJulio.Transferir(150, &contaDaMonalisa)
	fmt.Println(status)
	fmt.Println("Saldo da conta do Julio: ", contaDoJulio.saldo)
	fmt.Println("Saldo da conta da Monalisa: ", contaDaMonalisa.saldo)

}
