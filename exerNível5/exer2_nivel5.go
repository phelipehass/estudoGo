package main

import "fmt"

//- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
//- Demonstre os valores do map utilizando range.
//- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.

type person struct {
	nome              string
	sobrenome         string
	sorvetesFavoritos []string
}

func main() {
	pessoa1 := person{"Xundas", "Silva", []string{"Chocolate", "Morango", "Abacaxi"}}
	pessoa2 := person{"Alfredo", "Pereira", []string{"Uva", "Limão", "Maracujá"}}

	pessoas := map[string]person{
		pessoa1.sobrenome: pessoa1,
		pessoa2.sobrenome: pessoa2,
	}

	for _, value := range pessoas {
		fmt.Print(value.nome)
		fmt.Println("", value.sobrenome)

		for _, v := range value.sorvetesFavoritos {
			fmt.Println(v)
		}
	}

}
