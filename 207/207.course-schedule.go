package main

import "fmt"

/*
 * @lc app=leetcode.cn id=207 lang=golang
 * @lcpr version=21913
 *
 * [207] 课程表
 *
 * https://leetcode.cn/problems/course-schedule/description/
 *
 * algorithms
 * Medium (53.77%)
 * Likes:    1753
 * Dislikes: 0
 * Total Accepted:    331.7K
 * Total Submissions: 616.8K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
 *
 * 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi]
 * ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
 *
 *
 * 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
 *
 *
 * 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：true
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
 *
 * 示例 2：
 *
 * 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
 * 输出：false
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= numCourses <= 2000
 * 0 <= prerequisites.length <= 5000
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * prerequisites[i] 中的所有课程对 互不相同
 *
 *
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 构造课程依赖有向图, graph[i][j]为1时，代表要学习i必须先学习j
	graph := make([][]int, numCourses)
	for _, edge := range prerequisites {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	// 依次从每一门课程出发，进行深度遍历，判断图中是否形成环
	for i := 0; i < numCourses; i++ {
		if hasCycle(graph, make([]bool, numCourses), i) {
			return false
		}
	}

	return true
}

func hasCycle(graph [][]int, visited []bool, node int) bool {
	if visited[node] {
		return true
	}
	visited[node] = true
	for _, next := range graph[node] {
		if hasCycle(graph, visited, next) {
			return true
		}
	}

	visited[node] = false
	return false
}

// @lc code=end

/*
// @lcpr case=start
// 2\n[[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// 2\n[[1,0],[0,1]]\n
// @lcpr case=end

*/

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
