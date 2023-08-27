package main

// 641. 设计循环双端队列
// https://leetcode-cn.com/problems/design-circular-deque/
// 设计你的循环双端队列实现
// 实现 MyCircularDeque 类:
// - MyCircularDeque(int k) :构造函数,双端队列的大小为 k 。
// - boolean insertFront():将一个元素添加到双端队列头部。 如果操作成功返回 true 。
// - boolean insertLast() :将一个元素添加到双端队列尾部。如果操作成功返回 true 。
// - boolean deleteFront() :从双端队列头部删除一个元素。 如果操作成功返回 true 。
// - boolean deleteLast() :从双端队列尾部删除一个元素。如果操作成功返回 true 。
// - int getFront() ):从双端队列头部获得一个元素。如果双端队列为空,返回 -1 。
// - int getRear() :获得双端队列的最后一个元素。 如果双端队列为空,返回 -1 。
// - boolean isEmpty() :若双端队列为空,返回 true ;否则,返回 false 。
// - boolean isFull() :若双端队列满了,返回 true ; 否则,返回 false 。
// 思路:
// 使用循环队列的方式表示双端队列
type MyCircularDeque struct {
	data                  []int
	capacity, first, last int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data:     make([]int, k+1),
		capacity: k + 1,
	}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}

	d.first = (d.first - 1 + d.capacity) % d.capacity
	d.data[d.first] = value
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}

	d.data[d.last] = value
	d.last = (d.last + 1) % d.capacity

	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}
	d.first = (d.first + 1) % d.capacity
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}
	d.last = (d.last - 1 + d.capacity) % d.capacity
	return true
}

func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[d.first]
}

func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[(d.last-1+d.capacity)%d.capacity]
}

func (d *MyCircularDeque) IsEmpty() bool {
	return d.first == d.last
}

func (d *MyCircularDeque) IsFull() bool {
	return d.first == (d.last+1)%d.capacity
}

func main() {
	obj := Constructor(3)
	obj.InsertLast(1)
	obj.InsertLast(2)
	obj.InsertFront(3)
	obj.InsertFront(4)
	obj.GetRear()
	obj.IsFull()
	obj.DeleteLast()
	obj.InsertFront(4)
	obj.GetFront()
}
