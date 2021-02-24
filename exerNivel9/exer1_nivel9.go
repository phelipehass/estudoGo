package main

import (
	"fmt"
	"sync"
)

//- Alem da goroutine principal, crie duas outras goroutines.
//- Cada goroutine adicional devem fazer um print separado.
//- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go func1()
	go func2()
	go func3()
	wg.Wait()
}

func func1() {
	fmt.Println("One")
	wg.Done()
}

func func2() {
	fmt.Println("Two")
	wg.Done()
}

func func3() {
	fmt.Println("Tree")
	wg.Done()
}
