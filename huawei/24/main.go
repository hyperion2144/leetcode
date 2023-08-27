package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 給你一个整数数组nums，请计算数組的中心位置，数組的中心位置是数组中的一个下标，
// 其左侧所有元素相乘的积等于右侧所有元素相乘的积。数组第一个元素的左侧积为1，最后一个元素的右侧积为1。
// 如果数组有多个中心位置，应该返回最靠近左边的那一个，如果数组不存在中心位置，返回-1。
func main() {
	nums := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for _, text := range strings.Split(scanner.Text(), " ") {
		num, _ := strconv.Atoi(text)
		nums = append(nums, num)
	}

	leftProduct, rightSideProduct := make([]int, len(nums)), make([]int, len(nums))
	leftProduct[0] = 1
	rightSideProduct[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		rightSideProduct[i] = rightSideProduct[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		if leftProduct[i] == rightSideProduct[i] {
			fmt.Println(nums[i])
			return
		}
	}

	fmt.Println(-1)
}
