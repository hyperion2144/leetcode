package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	number, length := 0, len(nums)
	for index := 0; index < length-number; index++ {
		element := nums[index]
		if element == val {
			number++
			if index < length-number {
				nums[length-number], nums[index] = element, nums[length-number]
				index--
			}
		}
	}

	nums = nums[:len(nums)-number]

	return len(nums)
}

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
