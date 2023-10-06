package main

/*
 * @lc app=leetcode.cn id=714 lang=golang
 * @lcpr version=21917
 *
 * [714] 买卖股票的最佳时机含手续费
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
 *
 * algorithms
 * Medium (75.81%)
 * Likes:    980
 * Dislikes: 0
 * Total Accepted:    238.3K
 * Total Submissions: 314.3K
 * Testcase Example:  '[1,3,2,8,4,9]\n2'
 *
 * 给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。
 *
 * 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
 *
 * 返回获得利润的最大值。
 *
 * 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
 *
 *
 *
 * 示例 1：
 *
 * 输入：prices = [1, 3, 2, 8, 4, 9], fee = 2
 * 输出：8
 * 解释：能够达到的最大利润:
 * 在此处买入 prices[0] = 1
 * 在此处卖出 prices[3] = 8
 * 在此处买入 prices[4] = 4
 * 在此处卖出 prices[5] = 9
 * 总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8
 *
 * 示例 2：
 *
 * 输入：prices = [1,3,7,5,10,3], fee = 3
 * 输出：6
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 5 * 10^4
 * 1 <= prices[i] < 5 * 10^4
 * 0 <= fee < 5 * 10^4
 *
 *
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	// 使用dp表示第i天的最大利润
	// 则对于第i天，存在两种情况
	// 使用dp[i][0]表示第i天交易完成后不持有股票的最大利润，使用dp[i][1]表示第i天交易完成后持有股票的最大利润
	// 若第i天不持有股票，可能为前一天已经不持有股票，或者前一天持有股票，且在这天卖出，则最大利润为max(dp[i-1][0], dp[i-1][1]+price[i]-fee)
	// 若第i天持有股票，则可能为之前已持有股票，这天未进行操作，或者在这天买入股票，则最大利润为max(dp[i-1][1], dp[i-1][0]-price[i])
	// 则有如下状态转移方程:
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price[i]-fee)
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-price[i])
	// 从方程中可以发现，第i天的最大利润只和前一天的最大利润有关，而第i天的持有股票状态只和前一天的持有股票状态有关
	// 所以可以将dp简化为dp0,dp1，用于节省空间
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		ndp0 := max(dp0, dp1+prices[i]-fee)
		ndp1 := max(dp1, dp0-prices[i])

		dp0, dp1 = ndp0, ndp1
	}
	return dp0
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
// [1, 3, 2, 8, 4, 9]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,3,7,5,10,3]\n3\n
// @lcpr case=end

*/
