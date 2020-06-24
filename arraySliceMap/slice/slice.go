package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	//Slice nao é um array. Slice define um pedaço de um array.
	//No exemplo abaixo ele pega os elementos do indice 1 até o 3, ignorando o valor do elemento do indice 3
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2] // Aqui ele pega desde o indice 0 até o indice 2, ignorando o elemeto do indice 2
	fmt.Println(a2, s3)

	/* Podemos imaginar o slice como uma estrura que tem tamanho e um ponteiro
	para um elemento de um array
	*/
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
