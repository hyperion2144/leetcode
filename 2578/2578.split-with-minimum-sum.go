package main

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=2578 lang=golang
 * @lcpr version=21917
 *
 * [2578] 最小和分割
 *
 * https://leetcode.cn/problems/split-with-minimum-sum/description/
 *
 * algorithms
 * Easy (82.41%)
 * Likes:    44
 * Dislikes: 0
 * Total Accepted:    18.4K
 * Total Submissions: 22.3K
 * Testcase Example:  '4325'
 *
 * 给你一个正整数 num ，请你将它分割成两个非负整数 num1 和 num2 ，满足：
 *
 *
 * num1 和 num2 直接连起来，得到 num 各数位的一个排列。
 *
 *
 * 换句话说，num1 和 num2 中所有数字出现的次数之和等于 num 中所有数字出现的次数。
 *
 *
 * num1 和 num2 可以包含前导 0 。
 *
 *
 * 请你返回 num1 和 num2 可以得到的和的 最小 值。
 *
 * 注意：
 *
 *
 * num 保证没有前导 0 。
 * num1 和 num2 中数位顺序可以与 num 中数位顺序不同。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：num = 4325
 * 输出：59
 * 解释：我们可以将 4325 分割成 num1 = 24 和 num2 = 35 ，和为 59 ，59 是最小和。
 *
 *
 * 示例 2：
 *
 * 输入：num = 687
 * 输出：75
 * 解释：我们可以将 687 分割成 num1 = 68 和 num2 = 7 ，和为最优值 75 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 10 <= num <= 10^9
 *
 *
 */

// @lc code=start
func splitNum(num int) int {
	// 首先对数字num中每一位数进行排序,并去除数字0
	nums := []byte(strconv.Itoa(num))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// 将nums按长度均分为两份,求和
	num1, num2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			num1 = num1*10 + int(nums[i]-'0')
		} else {
			num2 = num2*10 + int(nums[i]-'0')
		}
	}

	return num1 + num2
}

// @lc code=end

/*
// @lcpr case=start
// 4325\n
// @lcpr case=end

// @lcpr case=start
// 687\n
// @lcpr case=end

*/

func main() {
	splitNum(4325)
	splitNum(687)
}
