package main

import "fmt"

//Para construir uma Struct eu preciso definir vários tipos dentro do tipo determiando
type ContaUsuario struct {
	titular string
	agencia int
	conta   int
	saldo   float64
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
	fmt.Println(gabriel, ana)
}
