package main

import (
	"fmt"
	"time"
)

func main() {

	/* O unico laço de repetição que temos em GO é o for*/

	//Simula um laço while
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando....")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	//Laço infinito

	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
