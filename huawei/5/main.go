package main

import (
	"fmt"
	"sort"
)

// splitArray function splits the elements of the nums array into several groups
// such that the sum of elements in each group is equal.
// It returns the minimum sum of elements within the groups.
func splitArray(nums []int) int {
	// Sort the elements
	sort.Ints(nums)

	n := len(nums)

	// Caculate the sum of elements
	sum := 0
	for _, nums := range nums {
		sum += nums
	}

	// If there are m group, the minimum sum is sum / m
	for m := n; m > 0; m-- {
		// If the sum is not divisible by m, continue
		if sum%m != 0 {
			continue
		}

		minisum := sum / m
		if nums[n-1] > minisum {
			continue
		}

		dp := make([]bool, 1<<n)
		sums := make([]int, 1<<n)

		dp[0] = true

		// Check if the sum of elements in each group is equal to minisum
		for i := 0; i < 1<<n; i++ {
			if !dp[i] {
				continue
			}

			for j := 0; j < n; j++ {
				if sums[i]+nums[j] > minisum {
					break
				}

				if i&(1<<j) == 0 {
					next := i | 1<<j
					if !dp[next] {

						sums[next] = (sums[i] + nums[j]) % minisum
						dp[next] = true
					}
				}
			}
		}

		if dp[(1<<n)-1] {
			return minisum
		}

	}

	return -1
}

func main() {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	minSum := splitArray(nums)
	fmt.Println(minSum)
}
