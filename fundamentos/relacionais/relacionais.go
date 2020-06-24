package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("==", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	/*
		Nos exemplos abaixo o go trata somente os valores das variaveis
		e nao leva em consideração seus endereços de memória
		porém ao utilizar ponteiros a regra muda
	*/
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1 == d2)
	fmt.Println(d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Jonas"}
	p2 := Pessoa{"Jonas"}

	fmt.Println("Pessoas:", p1 == p2)

}
