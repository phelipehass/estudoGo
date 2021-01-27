package main

import "fmt"

//- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//- Nome
//- Sobrenome
//- Sabores favoritos de sorvete
//- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.

type person struct {
	nome              string
	sobrenome         string
	sorvetesFavoritos []string
}

func main() {
	pessoa1 := person{"Xundas", "Silva", []string{"Chocolate", "Morango", "Abacaxi"}}
	pessoa2 := person{"Alfredo", "Pereira", []string{"Uva", "Limão", "Maracujá"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, value := range pessoa2.sorvetesFavoritos {
		fmt.Println(value)
	}
	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, value := range pessoa1.sorvetesFavoritos {
		fmt.Println(value)
	}

}
