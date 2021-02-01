package main

import (
	"fmt"
	"math"
)

//- Crie um tipo "quadrado"
//- Crie um tipo "círculo"
//- Crie um método "área" para cada tipo que calcule e retorne a área da figura
//- Área do círculo: 2 * π * raio
//- Área do quadrado: lado * lado
//- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
//- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
//- Crie um valor de tipo "quadrado"
//- Crie um valor de tipo "círculo"
//- Use a função "info" para demonstrar a área do "quadrado"
//- Use a função "info" para demonstrar a área do "círculo"

type rect struct {
	lado1 float32
	lado2 float32
}

type circulo struct {
	raio float32
}

func (c circulo) area() float32 {
	return 2 * math.Pi * c.raio
}

func (q rect) area() float32 {
	return q.lado1 * q.lado2
}

type figura interface {
	area() float32
}

func info(f figura) float32 {
	return f.area()
}

func main() {
	q := rect{15, 15}
	circle := circulo{32}

	fmt.Println("Área do quadrado:", info(q))
	fmt.Println("Área do círculo:", info(circle))
}
