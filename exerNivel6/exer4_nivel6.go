package main

import "fmt"

//- Crie um tipo struct "pessoa" que contenha os campos:
//- nome
//- sobrenome
//- idade
//- Crie um método para "pessoa" que demonstre o nome completo e a idade;
//- Crie um valor de tipo "pessoa";
//- Utilize o método criado para demonstrar esse valor.

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	person := pessoa{"Juca", "Sales", 19}
	printPerson(person)
}

func printPerson(p pessoa) {
	fmt.Println("Nome:", p.nome, p.sobrenome)
	fmt.Println("Idade", p.idade)
}
