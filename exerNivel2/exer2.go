package main

import "fmt"
/*- Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
- = 
- !=
- <=
- <
- >=
- >
- Demonstre estes valores.*/

func main() {
	a := (2 == 2)
	b := (3 != 4)
	c := (500 <= 450)
	d := (5 < 9)
	e := (3 >= 2)
	f := (4 > 5)

	fmt.Printf("a: %v\nb: %v\nc: %v\nd: %v\ne: %v\nf: %v\n",a,b,c,d,e,f)
	
}