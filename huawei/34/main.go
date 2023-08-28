package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 现在有多组整数数组，需要将他们合并成一个新的数组
// 合并规则：
// - 从每个数组里按顺序取出固定长度的内容合并到新的数组中，取完的内容会删除掉
// - 如果该行不足固定长度或者已经为空，则直接取出剩余部分的内容放到新的数组中，继续下一行
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var length, n int
	fmt.Scan(&length, &n)

	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, 0)
		scanner.Scan()
		for _, digit := range strings.Split(scanner.Text(), ",") {
			num, _ := strconv.Atoi(digit)
			nums[i] = append(nums[i], num)
		}
	}

	answer := make([]int, 0)
	for len(nums) > 0 {
		for i := 0; i < len(nums); i++ {
			if len(nums[i]) > length {
				answer = append(answer, nums[i][:length]...)
				nums[i] = nums[i][length:]
			} else {
				answer = append(answer, nums[i]...)
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
		}
	}

	fmt.Println(answer)
}
