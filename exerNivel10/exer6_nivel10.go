package main

import "fmt"

//- Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.

func main() {
	canal := make(chan int)
	go populateWithNumbers(100, canal)
	receive(canal)
}

func populateWithNumbers(total int, c chan int) {
	for i := 0; i < total; i++ {
		c <- i + 1
	}
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
