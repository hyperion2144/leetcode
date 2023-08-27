package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums) // [-4, -1,1,2]

	var (
		n   = len(nums)
		ans = math.MaxInt
	)

	update := func(sum int) {
		if abs(sum-target) < abs(ans-target) {
			ans = sum
		}
	}

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if sum := nums[i] + nums[i+1] + nums[i+2]; sum > target {
			update(sum)
			break
		}
		if sum := nums[i] + nums[n-2] + nums[n-1]; sum < target {
			update(sum)
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			update(sum)

			if sum <= target {
				for left += 1; left < right && nums[left] == nums[left-1]; left++ {
				}
			} else {
				for right -= 1; right > left && nums[right] == nums[right+1]; right-- {
				}
			}
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(threeSumClosest([]int{-100, -98, -2, -1}, -101))
}
