package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1670 lang=golang
 * @lcpr version=21913
 *
 * [1670] 设计前中后队列
 *
 * https://leetcode.cn/problems/design-front-middle-back-queue/description/
 *
 * algorithms
 * Medium (51.94%)
 * Likes:    33
 * Dislikes: 0
 * Total Accepted:    8.1K
 * Total Submissions: 15.5K
 * Testcase Example:  '["FrontMiddleBackQueue","pushFront","pushBack","pushMiddle","pushMiddle","popFront","popMiddle","popMiddle","popBack","popFront"]\n' +
  '[[],[1],[2],[3],[4],[],[],[],[],[]]'
 *
 * 请你设计一个队列，支持在前，中，后三个位置的 push 和 pop 操作。
 *
 * 请你完成 FrontMiddleBack 类：
 *
 *
 * FrontMiddleBack() 初始化队列。
 * void pushFront(int val) 将 val 添加到队列的 最前面 。
 * void pushMiddle(int val) 将 val 添加到队列的 正中间 。
 * void pushBack(int val) 将 val 添加到队里的 最后面 。
 * int popFront() 将 最前面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
 * int popMiddle() 将 正中间 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
 * int popBack() 将 最后面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
 *
 *
 * 请注意当有 两个 中间位置的时候，选择靠前面的位置进行操作。比方说：
 *
 *
 * 将 6 添加到 [1, 2, 3, 4, 5] 的中间位置，结果数组为 [1, 2, 6, 3, 4, 5] 。
 * 从 [1, 2, 3, 4, 5, 6] 的中间位置弹出元素，返回 3 ，数组变为 [1, 2, 4, 5, 6] 。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：
 * ["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle",
 * "pushMiddle", "popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
 * [[], [1], [2], [3], [4], [], [], [], [], []]
 * 输出：
 * [null, null, null, null, null, 1, 3, 4, 2, -1]
 *
 * 解释：
 * FrontMiddleBackQueue q = new FrontMiddleBackQueue();
 * q.pushFront(1);   // [1]
 * q.pushBack(2);    // [1, 2]
 * q.pushMiddle(3);  // [1, 3, 2]
 * q.pushMiddle(4);  // [1, 4, 3, 2]
 * q.popFront();     // 返回 1 -> [4, 3, 2]
 * q.popMiddle();    // 返回 3 -> [4, 2]
 * q.popMiddle();    // 返回 4 -> [2]
 * q.popBack();      // 返回 2 -> []
 * q.popFront();     // 返回 -1 -> [] （队列为空）
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= val <= 10^9
 * 最多调用 1000 次 pushFront， pushMiddle， pushBack， popFront， popMiddle 和 popBack
 * 。
 *
 *
*/

// @lc code=start

type DualEndQueueNode struct {
	Val  int
	Prev *DualEndQueueNode
	Next *DualEndQueueNode
}

type FrontMiddleBackQueue struct {
	length int
	head   *DualEndQueueNode
	tail   *DualEndQueueNode
}

func Constructor() FrontMiddleBackQueue {
	head := &DualEndQueueNode{}
	tail := &DualEndQueueNode{}
	head.Next, tail.Prev = tail, head
	return FrontMiddleBackQueue{
		head: head,
		tail: tail,
	}
}

func (queue *FrontMiddleBackQueue) PushFront(val int) {
	queue.length++
	node := &DualEndQueueNode{Val: val, Next: queue.head.Next, Prev: queue.head}
	queue.head.Next.Prev = node
	queue.head.Next = node
}

func (queue *FrontMiddleBackQueue) PushMiddle(val int) {
	pos := queue.length / 2
	cur := queue.head
	for i := 0; i < pos; i++ {
		cur = cur.Next
	}

	node := &DualEndQueueNode{Val: val, Prev: cur, Next: cur.Next}
	cur.Next.Prev = node
	cur.Next = node
	queue.length++
}

func (queue *FrontMiddleBackQueue) PushBack(val int) {
	queue.length++
	node := &DualEndQueueNode{Val: val, Prev: queue.tail.Prev, Next: queue.tail}
	queue.tail.Prev.Next = node
	queue.tail.Prev = node
}

func (queue *FrontMiddleBackQueue) PopFront() int {
	if queue.length == 0 {
		return -1
	}
	queue.length--
	val := queue.head.Next.Val
	queue.head.Next.Next.Prev = queue.head
	queue.head.Next = queue.head.Next.Next
	return val
}

func (queue *FrontMiddleBackQueue) PopMiddle() int {
	if queue.length == 0 {
		return -1
	}

	pos := (queue.length + 1) / 2
	cur := queue.head
	for i := 0; i < pos; i++ {
		cur = cur.Next
	}

	val := cur.Val
	cur.Prev.Next = cur.Next
	cur.Next.Prev = cur.Prev
	queue.length--
	return val
}

func (queue *FrontMiddleBackQueue) PopBack() int {
	if queue.length == 0 {
		return -1
	}

	queue.length--
	val := queue.tail.Prev.Val
	queue.tail.Prev.Prev.Next = queue.tail
	queue.tail.Prev = queue.tail.Prev.Prev
	return val
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
// @lc code=end

/*
// @lcpr case=start
// ["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle", "popFront", "popMiddle", "popMiddle","popBack", "popFront"][[], [1], [2], [3], [4], [], [], [], [], []]\n
// @lcpr case=end

*/

func main() {
	queue := Constructor()
	queue.PushFront(1)
	queue.PushBack(2)
	queue.PushMiddle(3)
	queue.PushMiddle(4)
	fmt.Println(queue.PopFront())
	fmt.Println(queue.PopMiddle())
	fmt.Println(queue.PopMiddle())
	fmt.Println(queue.PopBack())
	fmt.Println(queue.PopFront())
}
