package main

import (
	"fmt"
	"sort"
)

// 给定一个数组nums，可以将元素分为若干个组，使得每组和相等，求出满足条件的所有分组中，最大的平分组个数。
func main() {
	var m int
	fmt.Scan(&m)

	nums := make([]int, m)
	sum := 0
	for i := 0; i < m; i++ {
		fmt.Scan(&nums[i])
		sum += nums[i]
	}

	sort.Ints(nums)

	for g := m; g > 0; g-- {
		sums := make([]int, 1<<m)
		dp := make([]bool, 1<<m)
		dp[0] = true
		minisum := sum / g

		for i := 0; i < 1<<m; i++ {
			if !dp[i] {
				continue
			}

			for j := 0; j < m; j++ {
				if sums[i]+nums[j] > minisum {
					break
				}

				if i&(1<<j) == 0 {
					next := i | (1 << j)
					if !dp[next] {
						dp[next] = true
						sums[next] = (sums[i] + nums[j]) % minisum
					}
				}
			}
		}

		if dp[1<<m-1] {
			fmt.Println(g)
			return
		}
	}
}
