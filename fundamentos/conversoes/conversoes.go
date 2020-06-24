package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado...
	fmt.Println("Teste " + string(99))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para int
	num, _ := strconv.Atoi("123")
	// Esse método de conversao retorna 2 parametros, um que é o número e o outro que é um erro
	// caso o valor de conversao nao seja um numero inteiro
	fmt.Println(num)

	//string para boolean
	result, _ := strconv.ParseBool("true")
	fmt.Println(result)

}
