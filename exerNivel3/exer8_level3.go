package main

import "fmt"

func main() {
	x := 2

	switch {
	case x == 0:
		fmt.Println("X é zero")
	case x == 1:
		fmt.Println("X é 1")
	case x >= 2:
		fmt.Print("X é 2")
	}
}
