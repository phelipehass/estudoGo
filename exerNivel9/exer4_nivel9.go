package main

import (
	"fmt"
	"runtime"
	"sync"
)

//- Utilize mu para consertar a condição de corrida do exercício anterior.
var wg sync.WaitGroup
var mu = sync.Mutex{}
var contador int

func main() {
	x := 100

	criarGoroutines(x)
	wg.Wait()
	fmt.Println("Valor final", contador)
}

func criarGoroutines(i int) {
	wg.Add(i)

	for cont := 0; cont < i; cont++ {
		go func() {
			mu.Lock()
			t := contador
			runtime.Gosched()
			t++
			contador = t
			mu.Unlock()
			wg.Done()
		}()
	}
}
