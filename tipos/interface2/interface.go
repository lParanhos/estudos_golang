package main

import "fmt"

type espotivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// Nesse caso precisamos passar a referencia do elemento do tipo ferrari
	var carro2 espotivo = &ferrari{"f40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro2, carro1)
}
