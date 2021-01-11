package main

import "fmt"

//- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.

func main() {
	x := map[string][]string{
		"Santos_Phelpz":   {"Tocar guitarra", "games", "Correr"},
		"Barretos_Xundas": {"Bicicleta", "Tenis", "Natação"},
		"Barros_Alfredo":  {"Golf", "Futebol", "Correr"},
	}

	x["Guerra_Gustavo"] = []string{"Guitarra", "Solos", "Harmônicos"}

	delete(x, "Santos_Phelpz") //removendo

	for key, v := range x {
		for k, value := range v {
			fmt.Println(key, "-", k, value)
		}
	}
}
