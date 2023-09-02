package main

/*
 * @lc app=leetcode.cn id=554 lang=golang
 * @lcpr version=21913
 *
 * [554] 砖墙
 *
 * https://leetcode.cn/problems/brick-wall/description/
 *
 * algorithms
 * Medium (51.57%)
 * Likes:    322
 * Dislikes: 0
 * Total Accepted:    64.2K
 * Total Submissions: 124.5K
 * Testcase Example:  '[[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]'
 *
 * 你的面前有一堵矩形的、由 n 行砖块组成的砖墙。这些砖块高度相同（也就是一个单位高）但是宽度不同。每一行砖块的宽度之和相等。
 *
 * 你现在要画一条 自顶向下 的、穿过 最少
 * 砖块的垂线。如果你画的线只是从砖块的边缘经过，就不算穿过这块砖。你不能沿着墙的两个垂直边缘之一画线，这样显然是没有穿过一块砖的。
 *
 * 给你一个二维数组 wall ，该数组包含这堵墙的相关信息。其中，wall[i] 是一个代表从左至右每块砖的宽度的数组。你需要找出怎样画才能使这条线
 * 穿过的砖块数量最少 ，并且返回 穿过的砖块数量 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：wall = [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]
 * 输出：2
 *
 *
 * 示例 2：
 *
 * 输入：wall = [[1],[1],[1]]
 * 输出：3
 *
 *
 *
 * 提示：
 *
 *
 * n == wall.length
 * 1 <= n <= 10^4
 * 1 <= wall[i].length <= 10^4
 * 1 <= sum(wall[i].length) <= 2 * 10^4
 * 对于每一行 i ，sum(wall[i]) 是相同的
 * 1 <= wall[i][j] <= 2^31 - 1
 *
 *
 */

// @lc code=start
func leastBricks(wall [][]int) int {
	// 所有墙缝的出现次数，墙缝的位置用当前砖的长度表示
	wallJoint := make(map[int]int)
	for _, widths := range wall {
		// 当前砖长
		sum := 0
		// 最后一块砖的右侧墙缝不计入
		for i := 0; i < len(widths)-1; i++ {
			sum += widths[i]
			wallJoint[sum]++
		}
	}

	// 找出具有最多墙缝数量的位置
	maxJoint := 0
	for _, joint := range wallJoint {
		if joint > maxJoint {
			maxJoint = joint
		}
	}

	// 由于墙的高度固定，要穿过最少的砖，只要用高度减去最多的墙缝数量即可
	return len(wall) - maxJoint
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1],[1],[1]]\n
// @lcpr case=end

*/

