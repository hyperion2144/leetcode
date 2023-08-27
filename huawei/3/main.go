package main

import (
	"fmt"
	"sort"
)

// 部门组织绿岛骑行团建活动。租用公共双人自行车，每辆自行车最多坐两人，做最大载重M。
// 给出部门每个人的体重，请问最多需要租用多少双人自行车。
// 输入描述：
// 第一行两个数字m、n，分别代表自行车限重，部门门总人数
// 第二行，几个数字，代表每个人的体重，体重都小于等于自行车限重m。
// 0=m<=200
// 0≤n<=1000000
// 输出描述：
// 最小需要的双人自行车数量。
func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)

	w := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&w[i])
	}

	sort.Ints(w)

	var answer int
	for left, right := 0, n-1; left <= right; {
		lw, rw := w[left], w[right]

		if left == right {
			answer++
			break
		}

		if lw+rw > m {
			right--
			answer++
			continue
		}
		left++
		right--
		answer++
	}

	fmt.Println(answer)
}
