package main

import "fmt"

//- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
//- Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
//- Do quinto ao último item do slice (incluindo o último item!)
//- Do segundo ao sétimo item do slice (incluindo o sétimo item!)
//- Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
//- Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
func main() {
	slice := make([]int, 10, 20)

	for i := 0; i < 10; i++ {
		slice[i] = i + 1
	}

	fmt.Println(slice[:3])               //primeiro ao terceiro item
	fmt.Println(slice[len(slice)/2-1:])  //quinto ao último item
	fmt.Println(slice[1:7])              //segundo ao sétimo item
	fmt.Println(slice[2 : len(slice)-1]) //terceiro ao penúltimo item

}
