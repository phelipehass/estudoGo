package main

import "fmt"

/*- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.*/

const (
	_ = iota + 2020
	age1
	age2
	age3
	age4
)

func main() {
	fmt.Printf("%v, %v, %v, %v\n", age1, age2, age3, age4)
}
