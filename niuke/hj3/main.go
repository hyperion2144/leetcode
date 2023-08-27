package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var result = make([]bool, 501)

	for i := 0; i < N; i++ {
		var number int
		fmt.Scan(&number)

		if !result[number] {
			result[number] = true
		}
	}

	for index, value := range result {
		if value {
			fmt.Println(index)
		}
	}
}
