package main

/*
	Podemos criar um tipo personalizado a partir de um tipo base
	E associar novos metódos a eles usando os receivers
*/

type nota float64

func (n nota) entre(inicio, fim float64) bool {

	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 6.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}
func main() {

}
