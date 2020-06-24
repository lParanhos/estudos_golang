package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// A variavel cria nesse if, só pode ser usada dentro do bloco de controle,
	// bem similar como no laço for

	if i := numeroAleatorio(); i > 5 { // tambem é suportado no switch
		fmt.Println("Ganhou !!")
	} else {
		fmt.Println("Perdeu !!")
	}
}
