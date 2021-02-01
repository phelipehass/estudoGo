package main

import "fmt"

//- Crie e use um struct anônimo.
//- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

func main() {
	cliente := struct {
		nome     string
		dividas  map[string]int64
		aReceber []string
	}{
		nome:     "João",
		dividas:  map[string]int64{"carro": 1500, "celular": 350, "geladeira": 300},
		aReceber: []string{"venda de notebook", "investimentos"},
	}

	fmt.Println("Nome:", cliente.nome)
	fmt.Print("Dividas:")
	for i, v := range cliente.dividas {
		fmt.Println(i, "-", v)
	}
	fmt.Print("A receber: ")
	for _, v := range cliente.aReceber {
		fmt.Println(v)
	}
}
