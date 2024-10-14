package contas

import "banco/clientes"

type ContaUsuario struct {
	Titular clientes.Cliente
	Agencia int
	Conta   int
	Saldo   float64
}

func (conta *ContaUsuario) Sacar(valor float64) string {
	autorizadoASacar := valor <= conta.Saldo && valor > 0
	if autorizadoASacar {
		conta.Saldo -= valor
		return "Saldo sacado com sucesso"
	} else {
		return ("Algo deu errado, verifique seu saldo")
	}

}

func (conta *ContaUsuario) Depositar(valor float64) {
	if valor > 0 {
		conta.Saldo += valor
	} else {

	}

}

func (conta1 *ContaUsuario) Transferir(valor float64, conta2 *ContaUsuario) (string, float64) {
	if valor > 0 {
		conta1.Saldo -= valor
		conta2.Depositar(valor)

		return "Operação realizda com sucesso", conta1.Saldo
	} else {
		return "Operação iniviável", conta1.Saldo
	}
}
