package main

import "fmt"

func main() {
	//Crio um slice de 10 posicoes
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	//Crio um slice de 10 posicoes, porém o array interno tem 20 elementos
	//mas referencio apenas 10
	s = make([]int, 10, 20)

	//A funcao len, pega a capacidade maxima do meu slice
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	/* quando atingimos a capacidade maxima de um slice, e aumenta, proporcional ao tamanho definido, ao inves
	de gerar um erro
	*/
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
