package main

import "fmt"

//- Crie um novo tipo: veículo
//- O tipo subjacente deve ser struct
//- Deve conter os campos: portas, cor
//- Crie dois novos tipos: caminhonete e sedan
//- Os tipos subjacentes devem ser struct
//- Ambos devem conter "veículo" como struct embutido
//- O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
//- O tipo sedan deve conter um campo bool chamado "modeloLuxo"
//- Usando os structs veículo, caminhonete e sedan:
//- Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
//- Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
//- Demonstre estes valores.
//- Demonstre um único campo de cada um dos dois.

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	caminhonete1 := caminhonete{veiculo{2, "azul"}, false}
	sedan1 := sedan{veiculo{4, "prata bari"}, true}

	fmt.Println("Caminhonete:", caminhonete1)
	fmt.Println("Sedan:", sedan1)

	//Demonstrando um único campo de cada
	fmt.Println("Caminhonete 1 tem tração nas quatro:", caminhonete1.tracaoNasQuatro)
	fmt.Println("Sedan1 é de luxo:", sedan1.modeloLuxo)

}
