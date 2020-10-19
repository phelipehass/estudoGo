package main

import "fmt"

/*- Crie um programa que:
    - Atribua um valor int a uma variável
    - Demonstre este valor em decimal, binário e hexadecimal
    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
	- Demonstre esta outra variável em decimal, binário e hexadecimal
	*/

func main() {
	x:= 10
	fmt.Printf("Valor em decimal: %d, binário: %b, hexadecimal: %#x\n", x, x, x)

	z := x << 1 //deslocamento de bits
	fmt.Printf("Valor com bit deslocado em decimal: %d, binário: %b, hexadecimal: %#x\n", z, z, z)


}