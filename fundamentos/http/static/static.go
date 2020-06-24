package main

import (
	"log"
	"net/http"
)

func main() {
	//caminho do meu arquivo estático
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/", fs)

	log.Println("Executando ...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
