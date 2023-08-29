package main

import "fmt"

// 公司老板做了一比大生意，想要给每位员工分配一些奖金，想通过游戏的方式来决定每位员工分配的奖金。
// 按照工号排序，每个人随即抽取一个数字，按顺序往后遇到第一个数字比自己数字大的，那么这个人可获得"距离*数字差值”的奖金。
// 如果遇不到比自己数字大的，则这个人的奖金就是随机的数字。
func main() {
	var n int
	fmt.Scan(&n)

	random := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&random[i])
	}

	monotonicStack := make([][2]int, 0)
	answer := make([]int, n)
	for i, v := range random {
		if len(monotonicStack) > 0 && monotonicStack[len(monotonicStack)-1][1] < v {
			for _, vi := range monotonicStack {
				answer[vi[0]] = (i - vi[0]) * (v - vi[1])
			}
			monotonicStack = monotonicStack[:0]
		}
		monotonicStack = append(monotonicStack, [2]int{i, v})
	}

	for _, v := range monotonicStack {
		answer[v[0]] = v[1]
	}

	fmt.Println(answer)
}
