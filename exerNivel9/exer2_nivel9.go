package main

import "fmt"

//- Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
//- Crie um tipo para um struct chamado "pessoa"
//- Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
//- Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
//- Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
//- Demonstre no seu código:
//- Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
//- Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
//- Se precisar de dicas, veja: https://gobyexample.com/interfaces

type pessoa struct {
	nome      string
	sobrenome string
}

func (p *pessoa) falar() {
	fmt.Println("Olá", p.nome)
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	p := pessoa{"Xundas", "Faustina"}

	//dizerAlgumaCoisa(p) modo errado, pois deve passar o endereço da var
	dizerAlgumaCoisa(&p) //modo correto, está passando endereço da var
}
