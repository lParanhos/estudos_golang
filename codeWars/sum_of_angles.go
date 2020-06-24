package main

/*
https://www.codewars.com/kata/5a03b3f6a1c9040084001765/solutions/go
*/

func Angle(n int) int {
	s := (n - 2) * 180
	return s
}

func main() {
	Angle(3)
	Angle(4)
}
