package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516
	//Concatenando
	//fmt.Println("O valor de x é: " + x)
	xs := fmt.Sprint(x) // converte para string
	fmt.Println("O valor de x é: " + xs)
	//ou
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f.", x)

	a := 1
	b := 1.99999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	// %v imprime qualquer tipo
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
