package main

/*
 * @lc app=leetcode.cn id=309 lang=golang
 * @lcpr version=21917
 *
 * [309] 买卖股票的最佳时机含冷冻期
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 *
 * algorithms
 * Medium (64.48%)
 * Likes:    1625
 * Dislikes: 0
 * Total Accepted:    283.5K
 * Total Submissions: 439.7K
 * Testcase Example:  '[1,2,3,0,2]'
 *
 * 给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​
 *
 * 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
 *
 *
 * 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
 *
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1:
 *
 * 输入: prices = [1,2,3,0,2]
 * 输出: 3
 * 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
 *
 * 示例 2:
 *
 * 输入: prices = [1]
 * 输出: 0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 5000
 * 0 <= prices[i] <= 1000
 *
 *
 */

// @lc code=start
func maxProfit(prices []int) int {
	// dp[i]表示第i天的最大利润
	// 对于第i天，若当前未买入股票，则可以选择买入或者不操作，最大利润为dp[i-1]-prices[i] 或者 dp[i-1]
	// 对于第i天，若当前已买入股票，则可以选择卖出或者不操作，最大利润为dp[i-1]+prices[i] 或者 dp[i-1]
	// 若当前为冷冻期，则不操作
	n := len(prices)
	if n < 2 {
		return 0
	}

	// dp[i][0]: 手上持有股票的最大收益
	// dp[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// dp[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益

	dp := make([][3]int, n)
	dp[0][0] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}

	return max(dp[n-1][1], dp[n-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
