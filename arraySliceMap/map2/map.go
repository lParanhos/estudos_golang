package main

import "fmt"

func main() {
	funcsEsalario := map[string]float64{
		"Romario Cleberval": 1234.90,
		"Janete Souza":      15456.0,
		"Pedro Junior":      1200.0,
	}

	funcsEsalario["Cleito Bernards"] = 1300.0
	//Ao tentar deletar um elemento que nao existe no map, nenhum erro é retornado
	delete(funcsEsalario, "inexistente")

	for nome, salario := range funcsEsalario {
		fmt.Println(nome, salario)
	}
}
