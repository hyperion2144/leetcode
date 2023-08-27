package main

import "fmt"

// 小华和小薇一起通过玩积木游戏学习数学。
// 他们有很多积木，每个积木块上都有一个数字，积木块上的数字可能相同。
// 小华随机拿一些积木挨着排成一排，请小薇找到这排积木中数字相同且所处位置最远的2块积木块，计算他们的距离，小薇请你帮忙替她解决这个问题。
func main() {
	var n int
	fmt.Scan(&n)

	blocks := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		var block int
		fmt.Scan(&block)
		blocks[block] = append(blocks[block], i)
	}

	distance := -1
	for _, x := range blocks {
		if len(x) > 1 {
			distance = max(distance, x[len(x)-1]-x[0])
		}
	}

	fmt.Println(distance)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
