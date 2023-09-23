package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1286 lang=golang
 * @lcpr version=21914
 *
 * [1286] 字母组合迭代器
 *
 * https://leetcode.cn/problems/iterator-for-combination/description/
 *
 * algorithms
 * Medium (65.34%)
 * Likes:    67
 * Dislikes: 0
 * Total Accepted:    7.2K
 * Total Submissions: 11.1K
 * Testcase Example:  '["CombinationIterator","next","hasNext","next","hasNext","next","hasNext"]\n' +
  '[["abc",2],[],[],[],[],[],[]]'
 *
 * 请你设计一个迭代器类 CombinationIterator ，包括以下内容：
 *
 *
 * CombinationIterator(string characters, int combinationLength)
 * 一个构造函数，输入参数包括：用一个 有序且字符唯一 的字符串 characters（该字符串只包含小写英文字母）和一个数字
 * combinationLength 。
 * 函数 next() ，按 字典序 返回长度为 combinationLength 的下一个字母组合。
 * 函数 hasNext() ，只有存在长度为 combinationLength 的下一个字母组合时，才返回 true
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入:
 * ["CombinationIterator", "next", "hasNext", "next", "hasNext", "next",
 * "hasNext"]
 * [["abc", 2], [], [], [], [], [], []]
 * 输出：
 * [null, "ab", true, "ac", true, "bc", false]
 * 解释：
 * CombinationIterator iterator = new CombinationIterator("abc", 2); // 创建迭代器
 * iterator
 * iterator.next(); // 返回 "ab"
 * iterator.hasNext(); // 返回 true
 * iterator.next(); // 返回 "ac"
 * iterator.hasNext(); // 返回 true
 * iterator.next(); // 返回 "bc"
 * iterator.hasNext(); // 返回 false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= combinationLength <= characters.length <= 15
 * characters 中每个字符都 不同
 * 每组测试数据最多对 next 和 hasNext 调用 10^4次
 * 题目保证每次调用函数 next 时都存在下一个字母组合。
 *
 *
*/

// @lc code=start
type CombinationIterator struct {
	char              string
	combinationLength int
	next              []int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	next := make([]int, combinationLength)
	for i := range next {
		next[i] = i
	}
	return CombinationIterator{
		char:              characters,
		combinationLength: combinationLength,
		next:              next,
	}
}

func (it *CombinationIterator) add(index int) int {
	it.next[index]++
	if it.next[index] >= len(it.char)-(it.combinationLength-index-1) && index > 0 {
		it.next[index] = it.add(index-1) + 1
	}

	return it.next[index]
}

func (it *CombinationIterator) Next() string {
	if !it.HasNext() {
		return ""
	}

	nextChars := make([]byte, it.combinationLength)
	for i := 0; i < it.combinationLength; i++ {
		nextChars[i] = it.char[it.next[i]]
	}

	it.add(it.combinationLength - 1)

	return string(nextChars)
}

func (it *CombinationIterator) HasNext() bool {
	return it.next[it.combinationLength-1] < len(it.char)
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

/*
// @lcpr case=start
// ["CombinationIterator", "next", "hasNext", "next", "hasNext", "next", "hasNext"][["abc", 2], [], [], [], [], [], []]\n
// @lcpr case=end

// @lcpr case=start
// ["CombinationIterator","hasNext","next","hasNext","next","hasNext","next","next","next","hasNext","hasNext","next","hasNext","hasNext","next","hasNext","next","hasNext","hasNext","hasNext","next","next","hasNext","next","hasNext","next","next","hasNext","hasNext","next","next","hasNext","hasNext","next","hasNext","next","next","next","next","hasNext","hasNext","next","next","hasNext","hasNext","next","next","hasNext","next","hasNext","hasNext","hasNext","next","next","hasNext","hasNext","hasNext","hasNext","next","hasNext","next","hasNext","next","next","next","next","next","next","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext","hasNext"][["fiklnuy",3],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[]]\n
// @lcpr case=end

*/

func main() {
	it := Constructor("fiklnuy", 3)
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
