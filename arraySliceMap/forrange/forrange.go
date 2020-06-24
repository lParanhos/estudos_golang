package main

import "fmt"

func main() {

	//Desse modo, criamos um array de tamanho dinamico
	//Quem conta o tamanho do array é o compilador
	numeros := [...]int{1, 2, 3, 4, 5}

	for index, number := range numeros {
		fmt.Printf("%d) %d\n", index, number)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
