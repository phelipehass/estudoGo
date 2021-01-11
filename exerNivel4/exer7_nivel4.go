package main

import "fmt"

//- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
//- "Nome", "Sobrenome", "Hobby favorito"
//- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

func main() {
	person := [][]string{
		[]string{"Xundas", "Pereira", "Tocar guitarra"},
		[]string{"Genoveva", "Santos", "Programar"},
		[]string{"Mateus", "Silva", "Games"},
	}

	for _, v := range person {
		fmt.Println(v[0])
		for _, value := range v {
			fmt.Println("\t", value)
		}
	}
}
