package main

import "fmt"

/*
- Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
- Por exemplo:
    65
		U+0041 'A'
		U+0041 'A'
		U+0041 'A'
*/
func main() {
	for x := 65; x <= 90; x++ {
		repeatPrint := 0
		fmt.Println(x)

		for repeatPrint <= 2 {
			fmt.Printf("\t%#U\n", x)
			repeatPrint++
		}
	}
}
