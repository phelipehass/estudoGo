package main

import "fmt"

/*- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.*/

func main() {
	s := `Olá,
		Estou de volta!! :)`
	
	fmt.Printf("Raw string literal: %v\n", s)
}