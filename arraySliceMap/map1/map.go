package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//Maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123123123123] = "Maria"
	aprovados[123131545534] = "Jonas"
	aprovados[123131231231] = "Cleito"
	aprovados[454532223222] = "Janete"

	fmt.Println(aprovados)

	for key, value := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", value, key)
	}

	fmt.Println(aprovados[123123123123])
	delete(aprovados, 123123123123)
	fmt.Println(aprovados[123123123123])
}
