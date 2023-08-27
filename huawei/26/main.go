package main

import "fmt"

// 商人经营一家店铺，有number种商品，
// 由于仓库限制每件商品的最大持有数量是 item[index]
// 每种商品的价格是 item-price[index][day]
// 通过对商品的买进和卖出获取利润
// 请给出商人在 days 天内能获取的最大的利润
// 注:同一件商品可以反复买进和卖出
func main() {
	var number, days int
	fmt.Scan(&number, &days)

	item := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scan(&item[i])
	}

	itemPrice := make([][]int, number)
	for i := 0; i < number; i++ {
		itemPrice[i] = make([]int, days)
	}
	for i := 0; i < number; i++ {
		for j := 0; j < days; j++ {
			fmt.Scan(&itemPrice[i][j])
		}
	}

	fmt.Println(maxProfit(item, itemPrice))
}

func maxProfit(item []int, itemPrice [][]int) int {
	maxProfit := 0
	for i, prices := range itemPrice {
		profit := 0
		for j := 1; j < len(prices); j++ {
			profit += max(0, prices[j]-prices[j-1])
		}
		maxProfit += profit * item[i]
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
