package main

/* Return the number (count) of vowels in the given string.

We will consider a, e, i, o, and u as vowels for this Kata.

The input string will only consist of lower case letters and/or spaces.

 https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func GetCount(str string) (count int) {

	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	total := 0

	for _, value := range str {
		convertedValue := strings.ToLower(strconv.QuoteRune(value))
		for _, val := range vowels {
			if convertedValue == strconv.QuoteRune(val) {
				total++
			}
		}
	}

	return total
}

func main() {
	GetCount("mingau")
	fmt.Println(GetCount("mingauU"))
}
