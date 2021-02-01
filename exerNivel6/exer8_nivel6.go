package main

import "fmt"

//- Crie uma função que retorna uma função.
//- Atribua a função retornada a uma variável.
//- Chame a função retornada.

func main() {
	x := vezes100()

	fmt.Println(x(5))
}

func vezes100() func(int) int {
	return func(i int) int {
		return i * 100
	}
}
