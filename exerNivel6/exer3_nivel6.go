package main

import "fmt"

//- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.

func main() {
	x := 5
	y := 2
	defer fmt.Println("Ultima impressão!")
	fmt.Println(x + y)

}
