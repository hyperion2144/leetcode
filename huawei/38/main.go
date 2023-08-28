package main

import (
	"fmt"
	"sort"
)

// 小明决定去某旅游景点游玩，他在网上搜索到了各种价位的酒店(长度为n的数组A)，他的心理价位是x元
// 请帮他筛选出k个最接近x元的酒店(n>=k>0)，并由低到高打印酒店的价格
func main() {
	var n, k, x int
	fmt.Scan(&n, &k, &x)

	hotels := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&hotels[i])
	}

	sort.Slice(hotels, func(i, j int) bool {
		priceDifferencei := abs(hotels[i] - x)
		priceDifferencej := abs(hotels[j] - x)
		if priceDifferencei == priceDifferencej {
			return hotels[i] < hotels[j]
		}
		return priceDifferencei < priceDifferencej
	})

	answer := hotels[0:k]
	sort.Ints(answer)

	fmt.Println(answer)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
