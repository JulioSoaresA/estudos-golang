package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoJulio := contas.ContaCorrente{Titular: "Julio", NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.50}

	fmt.Println("Saldo da conta do Julio: ", contaDoJulio.Saldo)

	contaDaMonalisa := contas.ContaCorrente{Titular: "Monalisa", NumeroAgencia: 589, NumeroConta: 123457, Saldo: 500}

	fmt.Println("Saldo da conta da Monalisa: ", contaDaMonalisa.Saldo)

	status := contaDoJulio.Transferir(150, &contaDaMonalisa)
	fmt.Println(status)
	fmt.Println("Saldo da conta do Julio: ", contaDoJulio.Saldo)
	fmt.Println("Saldo da conta da Monalisa: ", contaDaMonalisa.Saldo)

}
