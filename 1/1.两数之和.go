package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	expected := make(map[int]int)
	for i, v := range nums {
		if j, ok := expected[target-v]; ok {
			return []int{j, i}
		}
		expected[v] = i
	}
	return nil
}

// @lc code=end
