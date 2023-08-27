package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadSlice('\n')
	input = input[:len(input)-1]

	if r := len(input) % 8; r > 0 {
		input = append(input, strings.Repeat("0", 8-r)...)
	}

	group := len(input) / 8
	for i := 0; i < group; i++ {
		fmt.Printf("%s\n", input[i*8:(i+1)*8])
	}
}
