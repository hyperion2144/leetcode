package main

import "fmt"

/*
 * @lc app=leetcode.cn id=146 lang=golang
 * @lcpr version=21917
 *
 * [146] LRU 缓存
 *
 * https://leetcode.cn/problems/lru-cache/description/
 *
 * algorithms
 * Medium (53.56%)
 * Likes:    2890
 * Dislikes: 0
 * Total Accepted:    530.4K
 * Total Submissions: 990.5K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
 *
 * 实现 LRUCache 类：
 *
 *
 *
 *
 * LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
 * int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
 * void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
 * key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
 *
 *
 * 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
 *
 *
 *
 *
 *
 * 示例：
 *
 * 输入
 * ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * 输出
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 *
 * 解释
 * LRUCache lRUCache = new LRUCache(2);
 * lRUCache.put(1, 1); // 缓存是 {1=1}
 * lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
 * lRUCache.get(1);    // 返回 1
 * lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
 * lRUCache.get(2);    // 返回 -1 (未找到)
 * lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
 * lRUCache.get(1);    // 返回 -1 (未找到)
 * lRUCache.get(3);    // 返回 3
 * lRUCache.get(4);    // 返回 4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= capacity <= 3000
 * 0 <= key <= 10000
 * 0 <= value <= 10^5
 * 最多调用 2 * 10^5 次 get 和 put
 *
 *
*/

// @lc code=start

type DListNode struct {
	Key  int
	Val  int
	Prev *DListNode
	Next *DListNode
}

type LRUCache struct {
	size       int
	capacity   int
	indexes    map[int]*DListNode
	head, tail *DListNode
}

func Constructor(capacity int) LRUCache {
	head := &DListNode{}
	tail := &DListNode{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		capacity: capacity,
		indexes:  make(map[int]*DListNode),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	n, ok := l.indexes[key]
	if !ok {
		return -1
	}

	// 将n移动到头部
	// 1. 删除原位置的节点
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	// 2. 插入到头部
	n.Next = l.head.Next
	n.Prev = l.head
	l.head.Next.Prev = n
	l.head.Next = n

	return n.Val
}

func (l *LRUCache) Put(key int, value int) {
	// key存在时，更新value
	if l.Get(key) != -1 {
		l.indexes[key].Val = value
		return
	}

	// 在头部插入新的节点
	l.size++
	n := &DListNode{
		Key:  key,
		Val:  value,
		Prev: l.head,
		Next: l.head.Next,
	}
	l.head.Next.Prev = n
	l.head.Next = n
	l.indexes[key] = n

	// 若插入时，超出容量，则删除最久未使用的节点
	if l.size > l.capacity {
		old := l.tail.Prev

		l.tail.Prev = old.Prev
		old.Prev.Next = l.tail

		delete(l.indexes, old.Key)
		l.size--
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

// @lcpr case=start
// ["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]\n
// @lcpr case=end

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}
