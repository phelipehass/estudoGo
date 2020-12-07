package main

import (
	"fmt"
)

/*
- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
    1. 42
    2. "James Bond"
    3. true
- Agora demonstre os valores nestas variáveis, com:
    1. Uma única declaração print
    2. Múltiplas declarações print
*/
func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println("Hello, playground")

	fmt.Printf("x: %v, y: %v, z: %v", x, y, z)
	fmt.Printf("\nx: %v", x)
	fmt.Printf("\ny: %v, z: %v", y)
	fmt.Printf("\nz: %v", z)
}
