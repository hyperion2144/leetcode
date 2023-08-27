package main

import (
	"fmt"
	"math"
)

var change = map[byte]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15}

func main() {
	var input string
	fmt.Scan(&input)

	input = input[2:]

	var result int

	length := len(input)
	for i := length; i > 0; i-- {
		h := change[input[i-1]]
		result += h * int(math.Pow(16, float64(length-i)))
	}

	fmt.Println(result)
}
