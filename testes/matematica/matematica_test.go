package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v"

/* Os métodos de teste precisam começar com Test
e receber o ponteiro de teste
*/
func TestMedia(t *testing.T) {
	t.Parallel()
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

/* Para rodar os tester por linha de comando: go test
ou
go test ./...
Esse comando roda todos os testes dentro do meu diretório
*/
