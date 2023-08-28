package main

import (
	"fmt"
)

// 对称就是最大的美学，现有一道关于对称字符串的美学
// 已知:
// 第一个字符串: R
// 第二个字符串: BR
// 第三个字符串: RBBR
// 第四个字符串: BRRBRBBR
// 第五个字符串: RBBRBRRBBRRBRBBR
// 既第i个字符串=第i-1号字符串的取反+第i-1号字符串。
// 取反即(R->B, B->R)
// 现在告诉你n和k，让你求第n个字符串的第k个字符是多少。
func main() {
	var t int
	fmt.Scan(&t)

	strs := make([][2]uint64, t)
	for i := 0; i < t; i++ {
		var n, k uint64
		fmt.Scan(&n, &k)
		strs[i] = [2]uint64{n, k}
	}

	for i := 0; i < t; i++ {
		find(strs[i][0], strs[i][1], true)
	}
}

func find(n, k uint64, negate bool) {
	if n == 1 && k == 0 {
		if negate {
			fmt.Println("red")
		} else {
			fmt.Println("blue")
		}
		return
	}

	mid := (uint64(1) << (n - 1)) / 2

	if k < mid {
		find(n-1, k, !negate)
	} else {
		find(n-1, k-mid, negate)
	}
}
