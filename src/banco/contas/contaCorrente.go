package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque > 0 && valorDoSaque <= c.Saldo {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else if valorDoSaque < 0 {
		return "Valor inválido"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso. Novo Saldo é de", c.Saldo
	} else {
		return "Falha ao realizar depósito. Valor inválido", valorDoDeposito
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) string {
	if valorDaTransferencia > 0 && valorDaTransferencia <= c.Saldo {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferência realizada com sucesso"
	} else {
		return "Falha ao realizar transferência."
	}
}
