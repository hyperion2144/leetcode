package main

import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=2034 lang=golang
 * @lcpr version=21917
 *
 * [2034] 股票价格波动
 *
 * https://leetcode.cn/problems/stock-price-fluctuation/description/
 *
 * algorithms
 * Medium (47.35%)
 * Likes:    160
 * Dislikes: 0
 * Total Accepted:    29.8K
 * Total Submissions: 62.9K
 * Testcase Example:  '["StockPrice","update","update","current","maximum","update","maximum","update","minimum"]\n' +
  '[[],[1,10],[2,5],[],[],[1,3],[],[4,2],[]]'
 *
 * 给你一支股票价格的数据流。数据流中每一条记录包含一个 时间戳 和该时间点股票对应的 价格 。
 *
 *
 * 不巧的是，由于股票市场内在的波动性，股票价格记录可能不是按时间顺序到来的。某些情况下，有的记录可能是错的。如果两个有相同时间戳的记录出现在数据流中，前一条记录视为错误记录，后出现的记录
 * 更正 前一条错误的记录。
 *
 * 请你设计一个算法，实现：
 *
 *
 * 更新 股票在某一时间戳的股票价格，如果有之前同一时间戳的价格，这一操作将 更正 之前的错误价格。
 * 找到当前记录里 最新股票价格 。最新股票价格 定义为时间戳最晚的股票价格。
 * 找到当前记录里股票的 最高价格 。
 * 找到当前记录里股票的 最低价格 。
 *
 *
 * 请你实现 StockPrice 类：
 *
 *
 * StockPrice() 初始化对象，当前无股票价格记录。
 * void update(int timestamp, int price) 在时间点 timestamp 更新股票价格为 price 。
 * int current() 返回股票 最新价格 。
 * int maximum() 返回股票 最高价格 。
 * int minimum() 返回股票 最低价格 。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：
 * ["StockPrice", "update", "update", "current", "maximum", "update",
 * "maximum", "update", "minimum"]
 * [[], [1, 10], [2, 5], [], [], [1, 3], [], [4, 2], []]
 * 输出：
 * [null, null, null, 5, 10, null, 5, null, 2]
 *
 * 解释：
 * StockPrice stockPrice = new StockPrice();
 * stockPrice.update(1, 10); // 时间戳为 [1] ，对应的股票价格为 [10] 。
 * stockPrice.update(2, 5);  // 时间戳为 [1,2] ，对应的股票价格为 [10,5] 。
 * stockPrice.current();     // 返回 5 ，最新时间戳为 2 ，对应价格为 5 。
 * stockPrice.maximum();     // 返回 10 ，最高价格的时间戳为 1 ，价格为 10 。
 * stockPrice.update(1, 3);  // 之前时间戳为 1 的价格错误，价格更新为 3 。
 * ⁠                         // 时间戳为 [1,2] ，对应股票价格为 [3,5] 。
 * stockPrice.maximum();     // 返回 5 ，更正后最高价格为 5 。
 * stockPrice.update(4, 2);  // 时间戳为 [1,2,4] ，对应价格为 [3,5,2] 。
 * stockPrice.minimum();     // 返回 2 ，最低价格时间戳为 4 ，价格为 2 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= timestamp, price <= 10^9
 * update，current，maximum 和 minimum 总 调用次数不超过 10^5 。
 * current，maximum 和 minimum 被调用时，update 操作 至少 已经被调用过 一次 。
 *
 *
*/

// @lc code=start
type StockPrice struct {
	min, max hp
	last     int
}

func Constructor() StockPrice {
	min := hp{
		indexes: make(map[int]int),
		prices:  make([][2]int, 0),
	}
	max := hp{
		indexes: make(map[int]int),
		prices:  make([][2]int, 0),
	}

	heap.Init(&min)
	heap.Init(&max)

	return StockPrice{
		min: min,
		max: max,
	}
}

func (s *StockPrice) Update(timestamp int, price int) {
	if i, ok := s.min.indexes[timestamp]; ok {
		s.min.prices[i][1] = price
		heap.Fix(&s.min, i)
	} else {
		heap.Push(&s.min, [2]int{timestamp, price})
	}

	if i, ok := s.max.indexes[timestamp]; ok {
		s.max.prices[i][1] = -price
		heap.Fix(&s.max, i)
	} else {
		heap.Push(&s.max, [2]int{timestamp, -price})
	}

	if s.last < timestamp {
		s.last = timestamp
	}
}

func (s *StockPrice) Current() int {
	return s.min.prices[s.min.indexes[s.last]][1]
}

func (s *StockPrice) Maximum() int {
	return -s.max.prices[0][1]
}

func (s *StockPrice) Minimum() int {
	return s.min.prices[0][1]
}

type hp struct {
	indexes map[int]int
	prices  [][2]int
}

func (h *hp) Len() int {
	return len(h.prices)
}

func (h *hp) Less(i int, j int) bool {
	return h.prices[i][1] < h.prices[j][1]
}

func (h *hp) Pop() any {
	v := h.prices[len(h.prices)-1]
	h.prices = h.prices[:len(h.prices)-1]
	delete(h.indexes, v[0])
	return v
}

func (h *hp) Push(x any) {
	h.prices = append(h.prices, x.([2]int))
	h.indexes[x.([2]int)[0]] = len(h.prices) - 1
}

func (h *hp) Swap(i int, j int) {
	h.indexes[h.prices[i][0]] = j
	h.indexes[h.prices[j][0]] = i
	h.prices[i], h.prices[j] = h.prices[j], h.prices[i]
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
// @lc code=end

/*
// @lcpr case=start
// ["StockPrice", "update", "update", "current", "maximum", "update", "maximum", "update", "minimum"]\n[[], [1, 10], [2, 5], [], [], [1, 3], [], [4, 2], []]\n
// @lcpr case=end
*/

func main() {
	c := Constructor()
	c.Update(38, 2308)
	fmt.Println(c.Current())
	fmt.Println(c.Maximum())
	fmt.Println(c.Minimum())
	c.Update(47, 7876)
	fmt.Println(c.Current())
	fmt.Println(c.Maximum())
	fmt.Println(c.Minimum())
	c.Update(58, 1866)
	fmt.Println(c.Current())
	fmt.Println(c.Maximum())
	fmt.Println(c.Minimum())
	c.Update(43, 121)
	fmt.Println(c.Current())
	fmt.Println(c.Maximum())
	fmt.Println(c.Minimum())
	c.Update(40, 5339)
	fmt.Println(c.Current())
	fmt.Println(c.Maximum())
	fmt.Println(c.Minimum())
	c.Update(32, 5339)
	fmt.Println(c.Current())
	fmt.Println(c.Maximum())
	fmt.Println(c.Minimum())
	c.Update(43, 6414)
	fmt.Println(c.Current())
	fmt.Println(c.Maximum())
	fmt.Println(c.Minimum())
	c.Update(49, 9639)
	fmt.Println(c.Current())
	fmt.Println(c.Maximum())
	fmt.Println(c.Minimum())
	c.Update(36, 3192)
	fmt.Println(c.Current())
	fmt.Println(c.Maximum())
	fmt.Println(c.Minimum())
	c.Update(48, 1006)
	fmt.Println(c.Current())
	fmt.Println(c.Maximum())
	fmt.Println(c.Minimum())
	c.Update(53, 8013)
	fmt.Println(c.Current())
	fmt.Println(c.Maximum())
	fmt.Println(c.Minimum())
}
