package main

import "fmt"

//- Demonstre o funcionamento de um closure.
//- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.

func main() {
	a := contPlus()
	b := contPlus()

	fmt.Println(a()) //1
	fmt.Println(a()) //2
	fmt.Println(a()) //3
	fmt.Println(b()) //1
	fmt.Println(b()) //2
	fmt.Println(b()) //3
}

func contPlus() func() int {
	cont := 0
	return func() int {
		cont += 1
		return cont
	}
}
