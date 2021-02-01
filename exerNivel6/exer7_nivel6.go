package main

import "fmt"

//- Atribua uma função a uma variável.
//- Chame essa função.

func main() {
	x := func(s string) {
		fmt.Println("Bom Dia", s)
	}

	x("xundinhas")

}
