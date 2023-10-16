/*
 * @lc app=leetcode.cn id=4 lang=golang
 * @lcpr version=30102
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (41.67%)
 * Likes:    6825
 * Dislikes: 0
 * Total Accepted:    1M
 * Total Submissions: 2.4M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
 *
 * 算法的时间复杂度应该为 O(log (m+n)) 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 使用两个指针，同时遍历数组nums1和nums2
	var s1, s2 int
	var num []int
	for s1 < len(nums1) && s2 < len(nums2) {
		n1, n2 := nums1[s1], nums2[s2]
		// 比较s1和s2所指向的数字，将较小的数字存入合并书组num，并向前移动一位
		if n1 <= n2 {
			num = append(num, n1)
			s1++
		} else if n2 < n1 {
			num = append(num, n2)
			s2++
		}
	}

	// 将未遍历到的数组内容直接放入合并数组
	if s1 < len(nums1) {
		num = append(num, nums1[s1:]...)
	}
	if s2 < len(nums2) {
		num = append(num, nums2[s2:]...)
	}

	// 合并数组长度为奇数，直接返回中间的数字
	// 否则取中间两个数的平均数
	if mid := len(num) / 2; len(num)%2 != 0 {
		return float64(num[mid])
	} else {
		return float64(num[mid-1]+num[mid]) / 2.0
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,3]\n[2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[3,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,1,1,1,1,1,1,1,4,4]\n[1,3,4,4,4,4,4,4,4,4,4]\n
// @lcpr case=end
*/

