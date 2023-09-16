/*
 * @lc app=leetcode.cn id=33 lang=golang
 * @lcpr version=21913
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (43.89%)
 * Likes:    2727
 * Dislikes: 0
 * Total Accepted:    773.5K
 * Total Submissions: 1.8M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 整数数组 nums 按升序排列，数组中的值 互不相同 。
 *
 * 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k],
 * nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始
 * 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
 *
 * 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
 *
 * 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 0
 * 输出：4
 *
 *
 * 示例 2：
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 3
 * 输出：-1
 *
 * 示例 3：
 *
 * 输入：nums = [1], target = 0
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5000
 * -10^4 <= nums[i] <= 10^4
 * nums 中的每个值都 独一无二
 * 题目数据保证 nums 在预先未知的某个下标上进行了旋转
 * -10^4 <= target <= 10^4
 *
 *
 */

// @lc code=start
func search(nums []int, target int) int {
	k := len(nums) - 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			k = i
			break
		}
	}

	// 小于最小
	if k+1 < len(nums) && target < nums[k+1] {
		return -1
	}
	// 大于最大
	if target > nums[k] {
		return -1
	}

	if target < nums[0] {
		// 在后半部分二分查找
		index := seartIns(nums[k+1:], target)
		if index == -1 {
			return -1
		}
		return k + index + 1
	} else {
		index := seartIns(nums[0:k+1], target)
		return index
	}
}

func seartIns(nums []int, x int) int {

	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) >> 1
		if nums[m] < x {
			l = m + 1
		} else if nums[m] > x {
			r = m
		} else {
			return m
		}
	}

	if l == r {
		if nums[l] == x {
			return l
		}
	}

	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [4,5,6,7,0,1,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [4,5,6,7,0,1,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/

