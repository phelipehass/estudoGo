package main

import "fmt"

/*- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.*/

func main() {
	num := 5

	fmt.Printf("Decimal: %d,\nBinário: %b,\nHexadecimal: %#x.\n ", num, num, num)
}
