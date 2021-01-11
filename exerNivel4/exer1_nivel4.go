package main

import "fmt"

//Usando uma literal composta:
//Crie um array que suporte 5 valores do tipo int
//atribua valores aos seus índices
//Utilizando format printing, demonstre o tipo do array.
func main() {
	arrayOne := [5]int{1, 2, 3, 4, 5}

	for _, value := range arrayOne {
		fmt.Println(value)
	}

	fmt.Printf("%T", arrayOne)
}
