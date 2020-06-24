package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em %s!\n", "GO")
}

/*
	Help
	go help [command]
	Especifica o que o comando faz/usa

	Versao
	go version

	Habilitar documentação offline
	godoc -http=:PORT
	Roda a documentação do go na porta especificada

	Informaçoes sobre o ambiente
	go env

	Documentação no editor sobre determinado comando
	go doc cmd/vet

	Realizar verificações no código
	go vet nomeDoArquivo.go

	Compilar código / Build
	go build nomeDoArquivo.go
	Obs: Para executar o mesmo, basta rodar
	./nomeDoArquivo

	Compilar e executar
	go run nomeDoArquivo.go

	Instalar pacotes
	go get -u url

*/
