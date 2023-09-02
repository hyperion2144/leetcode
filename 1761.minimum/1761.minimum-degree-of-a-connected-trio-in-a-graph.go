package main

/*
 * @lc app=leetcode.cn id=1761 lang=golang
 * @lcpr version=21913
 *
 * [1761] 一个图中连通三元组的最小度数
 *
 * https://leetcode.cn/problems/minimum-degree-of-a-connected-trio-in-a-graph/description/
 *
 * algorithms
 * Hard (46.60%)
 * Likes:    60
 * Dislikes: 0
 * Total Accepted:    10.3K
 * Total Submissions: 18.6K
 * Testcase Example:  '6\n[[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]]'
 *
 * 给你一个无向图，整数 n 表示图中节点的数目，edges 数组表示图中的边，其中 edges[i] = [ui, vi] ，表示 ui 和 vi
 * 之间有一条无向边。
 *
 * 一个 连通三元组 指的是 三个 节点组成的集合且这三个点之间 两两 有边。
 *
 * 连通三元组的度数 是所有满足此条件的边的数目：一个顶点在这个三元组内，而另一个顶点不在这个三元组内。
 *
 * 请你返回所有连通三元组中度数的 最小值 ，如果图中没有连通三元组，那么返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 6, edges = [[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]]
 * 输出：3
 * 解释：只有一个三元组 [1,2,3] 。构成度数的边在上图中已被加粗。
 *
 *
 * 示例 2：
 *
 * 输入：n = 7, edges = [[1,3],[4,1],[4,3],[2,5],[5,6],[6,7],[7,5],[2,6]]
 * 输出：0
 * 解释：有 3 个三元组：
 * 1) [1,4,3]，度数为 0 。
 * 2) [2,5,6]，度数为 2 。
 * 3) [5,6,7]，度数为 2 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= n <= 400
 * edges[i].length == 2
 * 1 <= edges.length <= n * (n-1) / 2
 * 1 <= ui, vi <= n
 * ui != vi
 * 图中没有重复的边。
 *
 *
 */

// @lc code=start
func minTrioDegree(n int, edges [][]int) int {
	g := make([]map[int]struct{}, n)
	h := make([][]int, n)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		g[i] = make(map[int]struct{})
	}

	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		g[x][y], g[y][x] = struct{}{}, struct{}{}
		degree[x]++
		degree[y]++
	}
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		if degree[x] < degree[y] || (degree[x] == degree[y] && x < y) {
			h[x] = append(h[x], y)
		} else {
			h[y] = append(h[y], x)
		}
	}

	ans := 0x3f3f3f3f
	for i := 0; i < n; i++ {
		for _, j := range h[i] {
			for _, k := range h[j] {
				if _, ok := g[i][k]; ok {
					ans = min(ans, degree[i]+degree[j]+degree[k]-6)
				}
			}
		}
	}
	if ans == 0x3f3f3f3f {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// 6\n[[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]]\n
// @lcpr case=end

// @lcpr case=start
// 7\n[[1,3],[4,1],[4,3],[2,5],[5,6],[6,7],[7,5],[2,6]]\n
// @lcpr case=end

*/

