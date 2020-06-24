package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	//Quando criamos um swtich sem uma expressao
	//ele procura o primeiro case que tenha uma condição verdadeira
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia !")
	case t.Hour() < 18:
		fmt.Println("Boa tarde !")
	default:
		fmt.Println("Boa noite !")
	}
}
