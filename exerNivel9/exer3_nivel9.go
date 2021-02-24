package main

import (
	"fmt"
	"runtime"
	"sync"
)

//- Utilizando goroutines, crie um programa incrementador:
//- Tenha uma variável com o valor da contagem
//- Crie um monte de goroutines, onde cada uma deve:
//- Ler o valor do contador
//- Salvar este valor
//- Fazer yield da thread com runtime.Gosched()
//- Incrementar o valor salvo
//- Copiar o novo valor para a variável inicial
//- Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
//- Demonstre que há uma condição de corrida utilizando a flag -race

var gr sync.WaitGroup

func main() {
	x := 100

	newgoroutines(x)
	gr.Wait()
}

func newgoroutines(i int) {
	gr.Add(i)
	contador := 0

	for cont := 0; cont < i; cont++ {
		y := contador
		go func(i int) {
			t := contador
			runtime.Gosched()
			t++
			contador = t
			gr.Done()
		}(y)
	}

	fmt.Println("Valor final", contador)
}
