package main

import "fmt"

/*- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.*/

const (
	x int = 10
	a = 5
	y = "teste"
	z = true
)
func main() {
	fmt.Printf("Constantes x: %d, a: %v, y: %v, z: %v\n",x, a, y, z)
}