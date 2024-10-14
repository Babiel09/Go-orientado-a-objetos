package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	//Criando o primeiro titular:
	gabriel := contas.ContaUsuario{
		Titular: clientes.Cliente{"Gabriel", "291.825.287.72", 28, false},
		Agencia: 123,
		Conta:   1,
	}

	ana := contas.ContaUsuario{
		Titular: clientes.Cliente{"ana", "123.455.234.72", 22, true},
		Agencia: 321,
		Conta:   2,
	}
	gabriel.Depositar(2423)
	ana.Depositar(4223)

	gabriel.Transferir(2423, &ana)
	ana.Transferir(2423, &gabriel)

	gabriel.VerSaldo()
	fmt.Println(ana)
}
