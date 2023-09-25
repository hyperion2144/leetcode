package main

import "fmt"

/*
 * @lc app=leetcode.cn id=460 lang=golang
 * @lcpr version=21917
 *
 * [460] LFU 缓存
 *
 * https://leetcode.cn/problems/lfu-cache/description/
 *
 * algorithms
 * Hard (45.59%)
 * Likes:    704
 * Dislikes: 0
 * Total Accepted:    69.4K
 * Total Submissions: 152.2K
 * Testcase Example:  '["LFUCache","put","put","get","put","get","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]'
 *
 * 请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。
 *
 * 实现 LFUCache 类：
 *
 *
 * LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
 * int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
 * void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量
 * capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用
 * 的键。
 *
 *
 * 为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。
 *
 * 当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。
 *
 * 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
 *
 *
 *
 * 示例：
 *
 * 输入：
 * ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get",
 * "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
 * 输出：
 * [null, null, null, 1, null, -1, 3, null, -1, 3, 4]
 *
 * 解释：
 * // cnt(x) = 键 x 的使用计数
 * // cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
 * LFUCache lfu = new LFUCache(2);
 * lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
 * lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
 * lfu.get(1);      // 返回 1
 * ⁠                // cache=[1,2], cnt(2)=1, cnt(1)=2
 * lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
 * ⁠                // cache=[3,1], cnt(3)=1, cnt(1)=2
 * lfu.get(2);      // 返回 -1（未找到）
 * lfu.get(3);      // 返回 3
 * ⁠                // cache=[3,1], cnt(3)=2, cnt(1)=2
 * lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
 * ⁠                // cache=[4,3], cnt(4)=1, cnt(3)=2
 * lfu.get(1);      // 返回 -1（未找到）
 * lfu.get(3);      // 返回 3
 * ⁠                // cache=[3,4], cnt(4)=1, cnt(3)=3
 * lfu.get(4);      // 返回 4
 * ⁠                // cache=[3,4], cnt(4)=2, cnt(3)=3
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= capacity <= 10^4
 * 0 <= key <= 10^5
 * 0 <= value <= 10^9
 * 最多调用 2 * 10^5 次 get 和 put 方法
 *
 *
*/

// @lc code=start

type DListNode struct {
	key, value, count int
	prev, next        *DListNode
}

type LFUCache struct {
	size, capacity int
	indexes        map[int]*DListNode
	counts         map[int]*DListNode
	head, tail     *DListNode
}

func Constructor(capacity int) LFUCache {
	head := &DListNode{}
	tail := &DListNode{}
	head.next = tail
	tail.prev = head
	return LFUCache{
		size:     0,
		capacity: capacity,
		indexes:  make(map[int]*DListNode),
		counts:   make(map[int]*DListNode),
		head:     head,
		tail:     tail,
	}
}

func (l *LFUCache) Get(key int) int {
	n := l.indexes[key]
	if n == nil {
		return -1
	}

	// 若当前节点是之前计数n的头结点，将该计数的头结点替换为下一节点
	// 否则，将节点移动到计数n头结点的前方
	if l.counts[n.count].key == key {
		next := n.next
		if next.count != n.count {
			delete(l.counts, n.count)
		} else {
			l.counts[n.count] = next
		}
	}

	// 当前节点的计数加1, 使其成为新的计数为n+1的头节点
	// 若原n+1的头节点不存在，则将节点移动到计数n的头结点前方
	// 若原n+1的头结点存在，则在原n+1的头结点前方插入
	n.count++
	cnt, exist := l.counts[n.count]
	if !exist {
		// 若当前节点不是之前计数n的头结点，将节点移动到计数n头结点的前方
		ncnt, ok := l.counts[n.count-1]
		if ok && n.next.key != ncnt.key {
			n.prev.next = n.next
			n.next.prev = n.prev

			n.prev = ncnt.prev
			n.next = ncnt
			ncnt.prev.next = n
			ncnt.prev = n
		}
	} else {
		n.prev.next = n.next
		n.next.prev = n.prev

		n.prev = cnt.prev
		n.next = cnt
		cnt.prev.next = n
		cnt.prev = n
	}

	l.counts[n.count] = n

	return n.value
}

func (l *LFUCache) Put(key int, value int) {
	// 当key存在时，更新对应value值
	if l.Get(key) != -1 {
		l.indexes[key].value = value
		return
	}

	// 当容量超出时，淘汰计数最少的节点
	if l.size+1 > l.capacity {
		removeNode := l.tail.prev
		removeNode.prev.next = removeNode.next
		removeNode.next.prev = removeNode.prev

		if l.counts[removeNode.count].key == removeNode.key {
			delete(l.counts, removeNode.count)
		}

		delete(l.indexes, removeNode.key)
	}

	// 插入新的节点，在对应计数为1的位置头部
	// 若当前计数为1的节点不存在，则直接在尾部插入
	// 若当前计数为1的节点存在，则在当前计数为1的头结点前方插入，并更换头结点
	n := &DListNode{
		key:   key,
		value: value,
		count: 1,
	}

	cnt, exist := l.counts[1]
	if !exist {
		n.prev = l.tail.prev
		n.next = l.tail

		l.tail.prev.next = n
		l.tail.prev = n
	} else {
		n.prev = cnt.prev
		n.next = cnt
		cnt.prev.next = n
		cnt.prev = n
	}

	l.size++
	l.indexes[key] = n
	l.counts[1] = n
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

/*
// @lcpr case=start
// ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]\n[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]\n
// @lcpr case=end

// @lcpr case=start
// ["LFUCache","put","put","get","get","get","put","put","get","get","get","get"]\n[[3],[2,2],[1,1],[2],[1],[2],[3,3],[4,4],[3],[2],[1],[4]]\n
// @lcpr case=end
*/

func main() {
	lfu := Constructor(3)
	lfu.Put(2, 2)
	lfu.Put(1, 1)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(2))
	lfu.Put(3, 3)
	lfu.Put(4, 4)
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(4))
}
