package main

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	for i, j := 0, 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && popped[j] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			j++
		}
	}

	return len(stack) == 0
}

func main() {
	validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})
}
