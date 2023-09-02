package main

/*
 * @lc app=leetcode.cn id=20 lang=golang
 * @lcpr version=21913
 *
 * [20] 有效的括号
 *
 * https://leetcode.cn/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (43.99%)
 * Likes:    4116
 * Dislikes: 0
 * Total Accepted:    1.6M
 * Total Submissions: 3.6M
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 每个右括号都有一个对应的相同类型的左括号。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "()"
 * 输出：true
 *
 *
 * 示例 2：
 *
 * 输入：s = "()[]{}"
 * 输出：true
 *
 *
 * 示例 3：
 *
 * 输入：s = "(]"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 仅由括号 '()[]{}' 组成
 *
 *
 */

// @lc code=start
func isValid(s string) bool {
	// 使用栈存储左括号，栈顶元素为当前期望匹配右括号的类型，若不匹配或栈为空则直接返回false
	stack := make([]rune, 0)
	for _, r := range s {
		if r == ')' || r == '}' || r == ']' {
			if len(stack) == 0 {
				return false
			}

			var expect rune
			switch stack[len(stack)-1] {
			case '(':
				expect = ')'
			case '{':
				expect = '}'
			case '[':
				expect = ']'
			}

			if r != expect {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}

	return len(stack) == 0
}

// @lc code=end

/*
// @lcpr case=start
// "()"\n
// @lcpr case=end

// @lcpr case=start
// "()[]{}"\n
// @lcpr case=end

// @lcpr case=start
// "(]"\n
// @lcpr case=end

// @lcpr case=start
// "]"\n
// @lcpr case=end

*/

