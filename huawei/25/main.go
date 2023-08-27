package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// findLongestPassword searches for the longest password in a password book
// that satisfies a specific condition.
// The condition is that when a character is removed from the end of the password,
// the resulting password is also present in the password book.
// If there are multiple passwords that satisfy the condition,
// the function returns the password with the highest lexicographical order.
// If no password satisfies the condition, the function returns an empty string.
func findLongestPassword() string {
	passwordBook := make([]string, 0)

	// Read input from standard input and store it in the password book
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	passwordBook = append(passwordBook, strings.Split(scanner.Text(), " ")...)

	// Sort the password book in ascending order of password length
	// If two passwords have the same length, sort them in lexicographical order
	sort.Slice(passwordBook, func(i, j int) bool {
		if len(passwordBook[i]) == len(passwordBook[j]) {
			return passwordBook[i] < passwordBook[j]
		}
		return len(passwordBook[i]) < len(passwordBook[j])
	})

	visited := make(map[string]bool, len(passwordBook))

	var answer string
	for _, password := range passwordBook {
		if len(password) == 1 || visited[password[:len(password)-1]] {
			visited[password] = true

			// Update the answer with the longest password found so far
			if len(password) >= len(answer) {
				answer = password
			}
		}
	}

	return answer
}

func main() {
	// Call the findLongestPassword function and print the result
	fmt.Println(findLongestPassword())
}
