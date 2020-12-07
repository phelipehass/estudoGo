package main

import "fmt"

//Crie um loop utilizando a sintaxe: for condition {}
//Utilize-o para demonstrar os anos desde que você nasceu
func main() {
	anoNascimento := 1998
	ano := 2020
	for anoNascimento <= ano {
		fmt.Println(anoNascimento)
		anoNascimento++
	}
}
