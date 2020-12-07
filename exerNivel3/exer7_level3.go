package main

import "fmt"

//Utilizando a solução anterior, adicione as opções else if e else.
func main() {
	trabalhar := false
	almoco := true

	if trabalhar {
		fmt.Println("Estou trabalhando")
	} else if almoco {
		fmt.Println("Estou no almoço")
	} else {
		fmt.Println("Estou descansando")
	}
}
