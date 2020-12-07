package main

import "fmt"

//Crie um loop utilizando a sintaxe: for
//Utilize-o para demonstrar os anos desde que você nasceu
func main() {
	anoNasci := 1998
	ano := 2020
	for {
		if anoNasci <= ano {
			fmt.Println(anoNasci)
			anoNasci++
		}
	}
}
