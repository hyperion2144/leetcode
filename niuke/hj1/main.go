package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var output int

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadBytes('\n')
	
	for i := len(input)-2; i >= 0; i-- {
		c := input[i]
		if (c == ' ') {
			break
		}

		output++
	}
	
	fmt.Println(output)
	
}
