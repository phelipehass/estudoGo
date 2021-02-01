package main

import "fmt"

//- Crie e utilize uma função anônima.

func main() {
	func(a, b int) {
		fmt.Println(a + b)
	}(5, 7)

}
