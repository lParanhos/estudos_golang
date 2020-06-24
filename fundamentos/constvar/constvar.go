package main

import (
	"fmt"
	"math"
)

// Em GO sempre que definimos uma variavel e nao a utilizamos
// É gerado um erro de compilação
func main() {
	//declaramos o nome e o tipo
	const PI float64 = 3.1415

	var raio = 3.2 //tipo (float64) inferido pelo compilador

	//forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	//Em blocos
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	//Declarando e atribuindo valores em uma linha
	e, f, g := 2, false, "epa!"
	fmt.Println(e, f, g)

	var h, i bool = false, true
	fmt.Println(h, i)

}
