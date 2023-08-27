package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	stringLine, _ := reader.ReadBytes('\n')
	stringLine = bytes.ToLower(stringLine[:len(stringLine)-1])

	char, _ := reader.ReadByte()
	chars := bytes.ToLower([]byte{char})

	count := bytes.Count(stringLine, chars)

	fmt.Println(count)
}
