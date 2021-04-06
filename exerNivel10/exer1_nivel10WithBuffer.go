package main

import "fmt"

//Faça esse código funcionar:
//func main() {
//	c := make(chan int)
//
//	c <- 42
//
//	fmt.Println(<-c)
//}
//-Usando uma função anônima auto-executável;
//-Usando Buffer

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)

}
