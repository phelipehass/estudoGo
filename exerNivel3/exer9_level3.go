package main

import "fmt"

func main() {
	esporteFavorito := "ping-pong"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("É futebol")
	case "Golf":
		fmt.Println("É patrão")
	case "ping-pong":
		fmt.Println("ping-pong?")
	}
}
