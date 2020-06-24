package main

import "fmt"

func main() {

	x := 1
	y := 2

	//Em go é apenas pos Fixada

	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // y -= 1 ou y = y - 1
	fmt.Println(y)

	//Nao permitido
	//	fmt.Println(x == y--)
}
