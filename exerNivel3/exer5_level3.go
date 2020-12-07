package main

import "fmt"

//Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
func main() {
	for x := 10; x <= 100; x++ {
		y := x % 4
		fmt.Printf("%d - O resto da divisão por 4 é %d\n", x, y)
	}
}
