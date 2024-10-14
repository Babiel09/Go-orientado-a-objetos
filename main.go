package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	//Criando o primeiro titular:
	gabriel := contas.ContaUsuario{
		Titular: clientes.Cliente{"Gabriel", 12345, 28, false},
		Agencia: 123,
		Conta:   1,
		Saldo:   9000,
	}

	ana := contas.ContaUsuario{
		Titular: clientes.Cliente{"ana", 12345, 22, true},
		Agencia: 321,
		Conta:   2,
		Saldo:   8000,
	}
	fmt.Println("Saldo do Gabriel:", gabriel.Saldo)
	fmt.Println("Saldo da Ana:", ana.Saldo)
	gabriel.Transferir(2000, &ana)
	fmt.Println("Saldo do Gabriel:", gabriel.Saldo)
	fmt.Println("Saldo da Ana:", ana.Saldo)
}
