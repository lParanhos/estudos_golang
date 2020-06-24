package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoluxuoso interface {
	esportivo
	luxuoso
	//Pode adicionar mais métodos
}

type bmw struct {
}

func (b bmw) ligarTurbo() {
	fmt.Println("Turbo ...")
}

func (b bmw) fazerBaliza() {
	fmt.Println("Baliza ...")
}

func main() {
	var b esportivoluxuoso = bmw{}
	b.ligarTurbo()
	b.fazerBaliza()
}
