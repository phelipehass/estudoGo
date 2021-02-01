package main

import "fmt"

//- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
//- Passe um valor do tipo slice of int como argumento para a função;
//- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
//- Passe um valor do tipo slice of int como argumento para a função.

func main() {
	slice1 := []int{5, 4, 6, 7, 8}

	fmt.Println("total is:", sumInt(slice1...))
	fmt.Println("Sum total is:", sumElements(slice1))
}

func sumInt(x ...int) int {
	sumTotal := 0
	for _, v := range x {
		sumTotal += v
	}
	return sumTotal
}

func sumElements(z []int) int {
	sumElementsTotal := 0
	for _, v := range z {
		sumElementsTotal += v
	}
	return sumElementsTotal
}
