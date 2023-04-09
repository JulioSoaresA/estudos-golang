package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteJulio := clientes.Titular{
		Nome:      "Júlio",
		CPF:       "123.456.789-10",
		Profissao: "Programador"}

	contaDoJulio := contas.ContaCorrente{
		Titular:       clienteJulio,
		NumeroAgencia: 589,
		NumeroConta:   123456}
	fmt.Println("Saldo da conta do Júlio é", contaDoJulio.ObterSaldo())
	fmt.Println(contaDoJulio.Depositar(100.25))
	fmt.Println("Saldo da conta do Júlio é", contaDoJulio.ObterSaldo())

	clienteMonalisa := clientes.Titular{
		Nome:      "Monalisa",
		CPF:       "987.654.321-10",
		Profissao: "Programadora"}

	contaDaMonalisa := contas.ContaCorrente{
		Titular:       clienteMonalisa,
		NumeroAgencia: 123,
		NumeroConta:   654321}
	fmt.Println("Saldo da conta da Monalisa é", contaDaMonalisa.ObterSaldo())

	status, valor := contaDaMonalisa.Depositar(500.90)
	fmt.Println(status, valor)

	clienteJunior := clientes.Titular{
		Nome:      "Junior",
		CPF:       "123.456.789-10",
		Profissao: "Autonomo"}

	contaDoJunior := contas.ContaPoupanca{
		Titular:       clienteJunior,
		NumeroAgencia: 123,
		NumeroConta:   123456,
		Operacao:      10}
	fmt.Println(contaDoJunior)
}
