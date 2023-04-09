package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else if valorDoSaque < 0 {
		return "Valor inválido"
	} else {
		return "Saldo insuficiente"
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

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
