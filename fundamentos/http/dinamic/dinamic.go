package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//Responsavel por escrever na resposta
func horaCerta(w http.ResponseWriter, r *http.Request) {
	//apesar de parecer uma data, esses números fazem parte da forma do Go, para formatar data e hora
	//cada um desses números tem um significado
	s := time.Now().Format("02/01/2006 03:04:05")
	//usamos para escrever no nosso writer
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando ...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
