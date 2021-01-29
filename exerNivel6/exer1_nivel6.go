package main

import "fmt"

//- Exercício:
//- Crie uma função que retorne um int
//- Crie outra função que retorne um int e uma string
//- Chame as duas funções
//- Demonstre seus resultados

func main() {
	fmt.Println("A soma de", 5, " + ", 6, "é:", somar(5, 6))
	fmt.Println(multiplicar(5, 6))
}

func somar(x, y int) int {
	return x + y
}

func multiplicar(x, y int) (string, int) {
	desc := "O resultado é:"
	return desc, x * y
}
