package main

import "fmt"

/*
	Quando trabalhamos com parametros do mesmo tipo,
	podemos colocar o tipo no final, como no exemplo abaixo.

	Dizemos que nossa funcao vai retornar 3 valor, do tipo bool
*/
func comprar(trab1, trab2 bool) (bool, bool, bool) {

	comprarTv55 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // mesmo resultado de "ou exclusivo"
	comprarSorvete := trab1 || trab2

	return comprarTv55, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := comprar(true, true)

	fmt.Printf("Tv50 : %t, Tv32: %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
}
