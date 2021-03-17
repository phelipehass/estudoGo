package main

import (
	"fmt"
	"sync"
)

//- 1. Todd:
//- Canais par, ímpar, e converge.
//- Func send manda pares pra um, ímpares pro outro, depois fecha.
//- Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!
//- Por fim um range retira todas as informações do canal converge.

func main() {
	par := make(chan int)
	ímpar := make(chan int)
	converge := make(chan int)

	go envia(par, ímpar)
	go recebe(par, ímpar, converge)

	for v := range converge {
		fmt.Println("Valor recebido:", v)
	}

}

func envia(p, i chan int) {
	x := 100
	for n := 0; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
	close(i)
}

func recebe(p, i, c chan int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(c)
}
