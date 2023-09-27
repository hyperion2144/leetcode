package main

import "container/heap"

/*
 * @lc app=leetcode.cn id=1333 lang=golang
 * @lcpr version=21917
 *
 * [1333] 餐厅过滤器
 *
 * https://leetcode.cn/problems/filter-restaurants-by-vegan-friendly-price-and-distance/description/
 *
 * algorithms
 * Medium (67.82%)
 * Likes:    57
 * Dislikes: 0
 * Total Accepted:    22.1K
 * Total Submissions: 32.5K
 * Testcase Example:  '[[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]\n1\n50\n10'
 *
 * 给你一个餐馆信息数组 restaurants，其中  restaurants[i] = [idi, ratingi, veganFriendlyi,
 * pricei, distancei]。你必须使用以下三个过滤器来过滤这些餐馆信息。
 *
 * 其中素食者友好过滤器 veganFriendly 的值可以为 true 或者 false，如果为 true 就意味着你应该只包括
 * veganFriendlyi 为 true 的餐馆，为 false 则意味着可以包括任何餐馆。此外，我们还有最大价格 maxPrice 和最大距离
 * maxDistance 两个过滤器，它们分别考虑餐厅的价格因素和距离因素的最大值。
 *
 * 过滤后返回餐馆的 id，按照 rating 从高到低排序。如果 rating 相同，那么按 id 从高到低排序。简单起见， veganFriendlyi
 * 和 veganFriendly 为 true 时取值为 1，为 false 时，取值为 0 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：restaurants =
 * [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]],
 * veganFriendly = 1, maxPrice = 50, maxDistance = 10
 * 输出：[3,1,5]
 * 解释：
 * 这些餐馆为：
 * 餐馆 1 [id=1, rating=4, veganFriendly=1, price=40, distance=10]
 * 餐馆 2 [id=2, rating=8, veganFriendly=0, price=50, distance=5]
 * 餐馆 3 [id=3, rating=8, veganFriendly=1, price=30, distance=4]
 * 餐馆 4 [id=4, rating=10, veganFriendly=0, price=10, distance=3]
 * 餐馆 5 [id=5, rating=1, veganFriendly=1, price=15, distance=1]
 * 在按照 veganFriendly = 1, maxPrice = 50 和 maxDistance = 10 进行过滤后，我们得到了餐馆 3, 餐馆
 * 1 和 餐馆 5（按评分从高到低排序）。
 *
 *
 * 示例 2：
 *
 * 输入：restaurants =
 * [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]],
 * veganFriendly = 0, maxPrice = 50, maxDistance = 10
 * 输出：[4,3,2,1,5]
 * 解释：餐馆与示例 1 相同，但在 veganFriendly = 0 的过滤条件下，应该考虑所有餐馆。
 *
 *
 * 示例 3：
 *
 * 输入：restaurants =
 * [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]],
 * veganFriendly = 0, maxPrice = 30, maxDistance = 3
 * 输出：[4,5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= restaurants.length <= 10^4
 * restaurants[i].length == 5
 * 1 <= idi, ratingi, pricei, distancei <= 10^5
 * 1 <= maxPrice, maxDistance <= 10^5
 * veganFriendlyi 和 veganFriendly 的值为 0 或 1 。
 * 所有 idi 各不相同。
 *
 *
 */

// @lc code=start
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	// 使用最大堆存储过滤后的餐馆信息
	// 最大堆按照rating和id从高到低排序
	h := make(hp, 0, len(restaurants))
	heap.Init(&h)

	for _, r := range restaurants {
		// 如果餐馆满足条件,则添加到最大堆中
		if veganFriendly == 1 && r[2] == 0 {
			continue
		}

		if r[3] <= maxPrice && r[4] <= maxDistance {
			heap.Push(&h, [2]int{r[0], r[1]})
		}
	}

	res := make([]int, 0, len(h))
	for h.Len() > 0 {
		// 取出最大堆中的餐馆信息
		restaurant := heap.Pop(&h).([2]int)
		res = append(res, restaurant[0])
	}

	return res
}

type hp [][2]int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i int, j int) bool {
	if h[i][1] == h[j][1] {
		return h[i][0] > h[j][0]
	}

	return h[i][1] > h[j][1]
}

func (h *hp) Pop() any {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func (h *hp) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h hp) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

// @lc code=end

/*
// @lcpr case=start
// [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]\n1\n50\n10\n
// @lcpr case=end

// @lcpr case=start
// [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]\n0\n50\n10\n
// @lcpr case=end

// @lcpr case=start
// [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]\n0\n30\n3\n
// @lcpr case=end

*/
