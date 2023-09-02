package main

/*
 * @lc app=leetcode.cn id=1654 lang=golang
 * @lcpr version=21913
 *
 * [1654] 到家的最少跳跃次数
 */

// @lc code=start
func minimumJumps(forbidden []int, a int, b int, x int) int {
	lower := 0
	maxVal := 0
	for _, val := range forbidden {
		maxVal = max(maxVal, val)
	}
	upper := max(maxVal+a, x) + b
	q := [][3]int{[3]int{0, 1, 0}}
	visited := make(map[int]bool)
	forbiddenSet := make(map[int]bool)
	visited[0] = true

	for _, position := range forbidden {
		forbiddenSet[position] = true
	}
	for len(q) > 0 {
		position, direction, step := q[0][0], q[0][1], q[0][2]
		q = q[1:]
		if position == x {
			return step
		}
		nextPosition, nextDirection := position+a, 1
		_, ok1 := visited[nextPosition*nextDirection]
		_, ok2 := forbiddenSet[nextPosition]
		if lower <= nextPosition && nextPosition <= upper && !ok1 && !ok2 {
			visited[nextPosition*nextDirection] = true
			q = append(q, [3]int{nextPosition, nextDirection, step + 1})
		}
		if direction == 1 {
			nextPosition, nextDirection := position-b, -1
			_, ok1 := visited[nextPosition*nextDirection]
			_, ok2 := forbiddenSet[nextPosition]
			if lower <= nextPosition && nextPosition <= upper && !ok1 && !ok2 {
				visited[nextPosition*nextDirection] = true
				q = append(q, [3]int{nextPosition, nextDirection, step + 1})
			}
		}
	}
	return -1
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
// [14,4,18,1,15]\n3\n15\n9\n
// @lcpr case=end

// @lcpr case=start
// [8,3,16,6,12,20]\n15\n13\n11\n
// @lcpr case=end

// @lcpr case=start
// [1,6,2,14,5,17,4]\n16\n9\n7\n
// @lcpr case=end

// @lcpr case=start
// [18,13,3,9,8,14]\n3\n8\n6\n
// @lcpr case=end


*/

