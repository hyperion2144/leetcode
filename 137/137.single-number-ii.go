package main

/*
 * @lc app=leetcode.cn id=137 lang=golang
 * @lcpr version=30102
 *
 * [137] 只出现一次的数字 II
 *
 * https://leetcode.cn/problems/single-number-ii/description/
 *
 * algorithms
 * Medium (71.97%)
 * Likes:    1106
 * Dislikes: 0
 * Total Accepted:    168.7K
 * Total Submissions: 234.4K
 * Testcase Example:  '[2,2,3,2]'
 *
 * 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
 *
 * 你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [2,2,3,2]
 * 输出：3
 *
 *
 * 示例 2：
 *
 * 输入：nums = [0,1,0,1,0,1,99]
 * 输出：99
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 * nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func singleNumber(nums []int) int {
	// 对于每个数字，其在二进制表示中，第i位的值为0或1
	// 若每个数字出现多次，则将其在i位的值相加，再对出现次数取余，得到的数字便是只出现一次的数字在i位的值
	// 在本题中，数字为大小为32位，因此可以使用int32
	var ans int32
	for i := 0; i < 32; i++ {
		// 对于一个数字右移i位后再与上1，得到的数字便是该数字在i位的值
		var sum int32
		for _, num := range nums {
			sum += (int32(num) >> i) & 1
		}
		// 若i位的和取余后的数字不为0，则该数字在i位的值为1，将ans的i位设置为1
		if sum%3 != 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

// @lc code=end

/*
// @lcpr case=start
// [2,2,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,1,0,1,99]\n
// @lcpr case=end

*/
