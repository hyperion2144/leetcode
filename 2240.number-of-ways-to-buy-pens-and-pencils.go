/*
 * @lc app=leetcode.cn id=2240 lang=golang
 * @lcpr version=21913
 *
 * [2240] 买钢笔和铅笔的方案数
 *
 * https://leetcode.cn/problems/number-of-ways-to-buy-pens-and-pencils/description/
 *
 * algorithms
 * Medium (57.54%)
 * Likes:    47
 * Dislikes: 0
 * Total Accepted:    17.4K
 * Total Submissions: 27.7K
 * Testcase Example:  '20\n10\n5'
 *
 * 给你一个整数 total ，表示你拥有的总钱数。同时给你两个整数 cost1 和 cost2
 * ，分别表示一支钢笔和一支铅笔的价格。你可以花费你部分或者全部的钱，去买任意数目的两种笔。
 *
 * 请你返回购买钢笔和铅笔的 不同方案数目 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：total = 20, cost1 = 10, cost2 = 5
 * 输出：9
 * 解释：一支钢笔的价格为 10 ，一支铅笔的价格为 5 。
 * - 如果你买 0 支钢笔，那么你可以买 0 ，1 ，2 ，3 或者 4 支铅笔。
 * - 如果你买 1 支钢笔，那么你可以买 0 ，1 或者 2 支铅笔。
 * - 如果你买 2 支钢笔，那么你没法买任何铅笔。
 * 所以买钢笔和铅笔的总方案数为 5 + 3 + 1 = 9 种。
 *
 *
 * 示例 2：
 *
 * 输入：total = 5, cost1 = 10, cost2 = 10
 * 输出：1
 * 解释：钢笔和铅笔的价格都为 10 ，都比拥有的钱数多，所以你没法购买任何文具。所以只有 1 种方案：买 0 支钢笔和 0 支铅笔。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= total, cost1, cost2 <= 10^6
 *
 *
 */

// @lc code=start
func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	if cost1 < cost2 {
		cost1, cost2 = cost2, cost1
	}
	var ans int
	for cost := 0; cost <= total; cost += cost1 {
		ans += (total-cost)/cost2 + 1
	}

	return int64(ans)
}

// @lc code=end

/*
// @lcpr case=start
// 20\n10\n5\n
// @lcpr case=end

// @lcpr case=start
// 5\n10\n10\n
// @lcpr case=end

*/

