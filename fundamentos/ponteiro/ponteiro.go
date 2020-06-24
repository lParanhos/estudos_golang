package main

import "fmt"

func main() {

	i := 1
	//Ponteiro é uma referencia de memória
	//Porem podemos usar nosso espaco de memória e compartilhar com nossas variaveis

	var p *int = nil

	//atribuo o endereço de memória da variavel i, para o meu ponteiro
	p = &i

	//Desrefenciando o ponteiro (pegando o valor)
	*p++
	i++

	//Go nao tem aritimeticos para ponteiros
	//p++

	fmt.Println(p, *p, i, &i)
}
