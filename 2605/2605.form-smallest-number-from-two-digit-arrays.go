package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=2605 lang=golang
 * @lcpr version=21913
 *
 * [2605] 从两个数字数组里生成最小数字
 *
 * https://leetcode.cn/problems/form-smallest-number-from-two-digit-arrays/description/
 *
 * algorithms
 * Easy (69.17%)
 * Likes:    60
 * Dislikes: 0
 * Total Accepted:    21.4K
 * Total Submissions: 30.9K
 * Testcase Example:  '[4,1,3]\n[5,7]'
 *
 * 给你两个只包含 1 到 9 之间数字的数组 nums1 和 nums2 ，每个数组中的元素 互不相同 ，请你返回 最小 的数字，两个数组都 至少
 * 包含这个数字的某个数位。
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [4,1,3], nums2 = [5,7]
 * 输出：15
 * 解释：数字 15 的数位 1 在 nums1 中出现，数位 5 在 nums2 中出现。15 是我们能得到的最小数字。
 *
 *
 * 示例 2：
 *
 * 输入：nums1 = [3,5,2,6], nums2 = [3,1,7]
 * 输出：3
 * 解释：数字 3 的数位 3 在两个数组中都出现了。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length, nums2.length <= 9
 * 1 <= nums1[i], nums2[i] <= 9
 * 每个数组中，元素 互不相同 。
 *
 *
 */

// @lc code=start
func minNumber(nums1 []int, nums2 []int) int {
	var repeat, lower1, lower2 int = math.MaxInt, math.MaxInt, math.MaxInt

	set := make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		if nums1[i] < lower1 {
			lower1 = nums1[i]
		}

		set[nums1[i]] = struct{}{}
	}

	for i := 0; i < len(nums2); i++ {
		if nums2[i] < lower2 {
			lower2 = nums2[i]
		}

		if _, ok := set[nums2[i]]; ok && nums2[i] < repeat {
			repeat = nums2[i]
		}
	}

	if repeat != math.MaxInt {
		return repeat
	}

	if lower1 < lower2 {
		return lower1*10 + lower2
	}

	return lower2*10 + lower1
}

// @lc code=end

/*
// @lcpr case=start
// [4,1,3]\n[5,7]\n
// @lcpr case=end

// @lcpr case=start
// [3,5,2,6]\n[3,1,7]\n
// @lcpr case=end

*/
