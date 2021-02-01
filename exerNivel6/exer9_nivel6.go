package main

import "fmt"

//- Callback: passe uma função como argumento a outra função.

func main() {
	bomdia(oi, "Xundas")
}

func bomdia(f func(s string) string, y string) {
	x := f(y)
	x += ", Bom dia!"
	fmt.Println(x)
}

func oi(s string) string {
	t := "Oi " + s
	return t
}
