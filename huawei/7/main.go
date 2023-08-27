package main

import (
	"fmt"
	"strconv"
	"strings"
)

// countBinaryNumbers counts the number of binary representations in the given range [l,r] that do not contain "101".
func countBinaryNumbers(l, r int) int {
	count := 0
	for num := l; num <= r; num++ {
		binary := strconv.FormatInt(int64(num), 2)
		if !strings.Contains(binary, "101") {
			count++
		}
	}
	return count
}

func main() {
	var l, r int
	fmt.Scan(&l, &r)
	result := countBinaryNumbers(l, r)
	fmt.Println(result)
}
