package main

import "fmt"

//Para construir uma Struct eu preciso definir vários tipos dentro do tipo determiando
type ContaUsuario struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

//Criando método para sacar dinheiro:
func (conta *ContaUsuario) sacar(valor float64) string {
	autorizadoASacar := valor <= conta.saldo && valor > 0
	if autorizadoASacar {
		conta.saldo -= valor
		return "Saldo sacado com sucesso"
	} else {
		return ("Algo deu errado, verifique seu saldo")
	}

}

func main() {
	//Criando o primeiro titular:
	gabriel := ContaUsuario{
		"Gabriel Castro",
		123,
		1,
		8793.98,
	}
	ana := ContaUsuario{
		"Ana Luiza (Irmã do Arthur)",
		321,
		2,
		6272.87,
	}

	fmt.Println(ana)
	gabriel.Depositar(7000)
	fmt.Println(gabriel)
}

func (conta *ContaUsuario) Depositar(valor float64) {
	conta.saldo += valor
	fmt.Println(conta.saldo)

}
