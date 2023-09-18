package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1452 lang=golang
 * @lcpr version=21913
 *
 * [1452] 收藏清单
 *
 * https://leetcode.cn/problems/people-whose-list-of-favorite-companies-is-not-a-subset-of-another-list/description/
 *
 * algorithms
 * Medium (51.26%)
 * Likes:    32
 * Dislikes: 0
 * Total Accepted:    8.1K
 * Total Submissions: 15.8K
 * Testcase Example:  '[["leetcode","google","facebook"],["google","microsoft"],["google","facebook"],["google"],["amazon"]]'
 *
 * 给你一个数组 favoriteCompanies ，其中 favoriteCompanies[i] 是第 i 名用户收藏的公司清单（下标从 0
 * 开始）。
 *
 * 请找出不是其他任何人收藏的公司清单的子集的收藏清单，并返回该清单下标。下标需要按升序排列。
 *
 *
 *
 * 示例 1：
 *
 * 输入：favoriteCompanies =
 * [["leetcode","google","facebook"],["google","microsoft"],["google","facebook"],["google"],["amazon"]]
 * 输出：[0,1,4]
 * 解释：
 * favoriteCompanies[2]=["google","facebook"] 是
 * favoriteCompanies[0]=["leetcode","google","facebook"] 的子集。
 * favoriteCompanies[3]=["google"] 是
 * favoriteCompanies[0]=["leetcode","google","facebook"] 和
 * favoriteCompanies[1]=["google","microsoft"] 的子集。
 * 其余的收藏清单均不是其他任何人收藏的公司清单的子集，因此，答案为 [0,1,4] 。
 *
 *
 * 示例 2：
 *
 * 输入：favoriteCompanies =
 * [["leetcode","google","facebook"],["leetcode","amazon"],["facebook","google"]]
 * 输出：[0,1]
 * 解释：favoriteCompanies[2]=["facebook","google"] 是
 * favoriteCompanies[0]=["leetcode","google","facebook"] 的子集，因此，答案为 [0,1] 。
 *
 *
 * 示例 3：
 *
 * 输入：favoriteCompanies = [["leetcode"],["google"],["facebook"],["amazon"]]
 * 输出：[0,1,2,3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= favoriteCompanies.length <= 100
 * 1 <= favoriteCompanies[i].length <= 500
 * 1 <= favoriteCompanies[i][j].length <= 20
 * favoriteCompanies[i] 中的所有字符串 各不相同 。
 * 用户收藏的公司清单也 各不相同 ，也就是说，即便我们按字母顺序排序每个清单， favoriteCompanies[i] !=
 * favoriteCompanies[j] 仍然成立。
 * 所有字符串仅包含小写英文字母。
 *
 *
 */

// @lc code=start
func peopleIndexes(favoriteCompanies [][]string) []int {
	for _, f := range favoriteCompanies {
		sort.Strings(f)
	}

	sorted := make([][]string, len(favoriteCompanies))
	copy(sorted, favoriteCompanies)

	sort.Slice(sorted, func(i, j int) bool {
		return len(sorted[i]) >= len(sorted[j])
	})

	ans := make([]int, 0)
	for i, f := range favoriteCompanies {
		flag := true
		for _, s := range sorted {
			if len(f) >= len(s) {
				break
			}

			l, p := 0, 0
			for p < len(s) && l < len(f) {
				if s[p] == f[l] {
					l++
				}
				p++
			}

			if l == len(f) {
				flag = false
				break
			}
		}

		if flag {
			ans = append(ans, i)
		}
	}

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [["leetcode","google","facebook"],["google","microsoft"],["google","facebook"],["google"],["amazon"]]\n
// @lcpr case=end

// @lcpr case=start
// [["leetcode","google","facebook"],["leetcode","amazon"],["facebook","google"]]\n
// @lcpr case=end

// @lcpr case=start
// [["leetcode"],["google"],["facebook"],["amazon"]]\n
// @lcpr case=end

// @lcpr case=start
// [["nxaqhyoprhlhvhyojanr","ovqdyfqmlpxapbjwtssm","qmsbphxzmnvflrwyvxlc","udfuxjdxkxwqnqvgjjsp","yawoixzhsdkaaauramvg","zycidpyopumzgdpamnty"],["nxaqhyoprhlhvhyojanr","ovqdyfqmlpxapbjwtssm","udfuxjdxkxwqnqvgjjsp","yawoixzhsdkaaauramvg","zycidpyopumzgdpamnty"],["ovqdyfqmlpxapbjwtssm","qmsbphxzmnvflrwyvxlc","udfuxjdxkxwqnqvgjjsp","yawoixzhsdkaaauramvg","zycidpyopumzgdpamnty"]]\n
// @lcpr case=end

*/

func main() {
	fmt.Println(peopleIndexes([][]string{
		{"leetcode", "google", "facebook"},
		{"google", "microsoft"},
		{"google", "facebook"},
		{"google"},
		{"amazon"},
	}))
}
