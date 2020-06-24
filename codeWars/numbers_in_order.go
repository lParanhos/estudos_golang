package main

import "fmt"

//https://www.codewars.com/kata/56b7f2f3f18876033f000307/train/go

func InAscOrder(numbers []int) bool {

	result := true
	current := 0

	for index, value := range numbers {
		if index == 0 {
			current = value
		} else if current > value {
			result = false
			break
		} else {
			current = value
		}
	}

	return result
}

func main() {

	fmt.Println(InAscOrder([]int{1, 2, 1, 5, 6}))
}

//outras resoluções
//https://www.codewars.com/kata/56b7f2f3f18876033f000307/solutions/go
