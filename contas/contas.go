package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaUsuario struct {
	Titular        clientes.Cliente
	Agencia, Conta int
	saldo          float64 //O saldo ficara privado para o usuário não o alterar
}

func (conta *ContaUsuario) Sacar(valor float64) string {
	autorizadoASacar := valor <= conta.saldo && valor > 0
	if autorizadoASacar {
		conta.saldo -= valor
		return "saldo sacado com sucesso"
	} else {
		return ("Algo deu errado, verifique seu saldo")
	}

}

func (conta *ContaUsuario) Depositar(valor float64) {
	if valor > 0 {
		conta.saldo += valor
	} else {

	}

}

func (conta1 *ContaUsuario) Transferir(valor float64, conta2 *ContaUsuario) (string, float64) {
	if valor > 0 {
		conta1.saldo -= valor
		conta2.Depositar(valor)

		return "Operação realizda com sucesso", conta1.saldo
	} else {
		return "Operação iniviável", conta1.saldo
	}
}

func (conta *ContaUsuario) VerSaldo() {
	fmt.Println(conta.saldo)

}
