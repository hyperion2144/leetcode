package main

import (
	"fmt"
)

// 模拟商场优惠打折，有三种优惠券可以用，满减券、打折券和无门槛券。
// 满减券: 满100减10, 满200减20, 满300减30， 满400减40, 以此类推不限制使用
// 打折券: 固定折扣92折，且打折之后向下取整，每次购物只能用1次
// 无门槛券: 一张券减5元，没有使用限制
// 每个人结账使用优惠券时有以下限制: 每人每次只能用两种优惠券，并且同一种优惠券必须一次用完，不能跟别的穿插使用(比如用一张满减券，再用一张打折券，再用一张满减券，这种顺序不行)
// 求不同使用顺序下每个人用完券之后得到的最低价格和对应使用优惠券的总数，如果两种顺序得到的价格一样低，就取使用优惠券数量较少的顺序
func main() {
	var coupon1, coupon2, coupon3 int
	fmt.Scan(&coupon1, &coupon2, &coupon3)

	var n int
	fmt.Scan(&n)

	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&prices[i])
	}

	for _, price := range prices {
		fmt.Println(minPrice(price, coupon1, coupon2, coupon3))
	}
}

func minPrice(price int, coupon1, coupon2, coupon3 int) (int, int) {
	ans, count := price, 0

	var mp func(int, int, int, int, int, int)
	mp = func(level, p, c, c1, c2, c3 int) {
		if c1 == 0 && c2 == 0 && c3 == 0 || p == 0 || level == 3 {
			if p < ans {
				ans = p
				count = c
			} else if p == ans {
				count = min(count, c)
			}
			return
		}

		if c1 > 0 && p >= 100 {
			pi := p
			cnt := c
			for i := 0; i < coupon1; i++ {
				cnt++
				pi = pi - (pi-pi%10)/10
				if pi < 100 {
					break
				}
			}
			mp(level+1, pi, cnt, 0, c2, c3)
		}

		if c2 > 0 {
			mp(level+1, int(float64(p)*0.92), c+1, c1, 0, c3)
		}

		if c3 > 0 {
			pi := p
			cnt := c
			for i := 0; i < coupon3; i++ {
				cnt++
				pi = pi - 5
				if pi <= 0 {
					pi = 0
					break
				}
			}
			mp(level+1, pi, cnt, c1, c2, 0)
		}
	}

	mp(1, price, 0, coupon1, coupon2, coupon3)
	return ans, count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
