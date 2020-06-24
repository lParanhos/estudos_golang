package main

import "fmt"

//https://www.codewars.com/kata/51b6249c4612257ac0000005/train/go

func Decode(roman string) int {
	total := 0

	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for _, value := range roman {
		total += values[value]
	}

	return total
}

func main() {
	fmt.Println(Decode("XXI"))
	Decode("XXI")
}
