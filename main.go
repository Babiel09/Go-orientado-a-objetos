package main

import "fmt"

func main() {
	//Criando o primeiro titular:
	var titular1 string = "Gabriel"
	var agencia1 int = 123
	var conta int = 1
	var saldo float64 = 234.56
	fmt.Println("O titular ", titular1, " possui R$", saldo, " em sua conta bancária.")
	fmt.Println("O titular ", titular1, " está cadastrado na agencia: ", agencia1, " e sau conta é: ", conta)

	//Criando o segundo titular:
	var titular2 string = "Ana Luiza de Lima"
	var agencia2 int = 321
	var conta2 int = 2
	var saldo2 float64 = 23.56

	fmt.Println("A titular ", titular2, " possui R$", saldo2, " em sua conta bancária.")
	fmt.Println("A titular ", titular2, " está cadastrado na agencia: ", agencia2, " e sau conta é: ", conta2)
}
