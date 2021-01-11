package main

import "fmt"

//- Crie um map com key tipo string e value tipo []string.
//- Key deve conter nomes no formato sobrenome_nome
//- Value deve conter os hobbies favoritos da pessoa
//- Demonstre todos esses valores e seus índices.

func main() {
	x := map[string][]string{
		"Santos_Phelpz":   {"Tocar guitarra", "games", "Correr"},
		"Barretos_Xundas": {"Bicicleta", "Tenis", "Natação"},
		"Barros_Alfredo":  {"Golf", "Futebol", "Correr"},
	}

	for key, v := range x {
		for k, value := range v {
			fmt.Println(key, "-", k, value)
		}
	}
}
