package main

import "fmt"

//- Usando uma literal composta:
//- Crie uma slice de tipo int
//- Atribua 10 valores a ela
//- Utilize range para demonstrar todos estes valores.
//- E utilize format printing para demonstrar seu tipo.
func main() {
	sliceOne := make([]int, 10, 20)

	for i := 0; i < 10; i++ {
		sliceOne[i] = i + 1
	}

	for _, value := range sliceOne {
		fmt.Println(value)
	}
	fmt.Printf("%T", sliceOne)
}
