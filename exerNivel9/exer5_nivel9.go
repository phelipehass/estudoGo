package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//- Utilize atomic para consertar a condição de corrida do exercício #3.

var gr sync.WaitGroup
var contador int32

func main() {
	x := 100

	newGoroutines(x)
	gr.Wait()
	fmt.Println("Valor final", contador)
}

func newGoroutines(i int) {
	gr.Add(i)
	for cont := 0; cont < i; cont++ {
		go func() {
			atomic.AddInt32(&contador, 1)
			v := atomic.LoadInt32(&contador)
			runtime.Gosched()
			fmt.Println(v)
			gr.Done()
		}()
	}

}
