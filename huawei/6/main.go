package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 求单向链表Q中间的节点值，如果奇数个节点取中间，偶数个取偏右边的那个值。
func middleNode(head *ListNode) int {
	if head == nil {
		return 0
	}

	slow, fast, cnt := head, head, 0
	for fast.Next != nil {
		fast = fast.Next
		cnt++

		if cnt%2 == 0 {
			slow = slow.Next
		}
	}

	if cnt%2 == 1 {
		slow = slow.Next
	}

	return slow.Val
}

func main() {
	var head, n int
	fmt.Scanf("%d %d", &head, &n)

	type tempNode struct {
		Val  int
		Addr int
		Next int
	}

	nodes := make(map[int]*tempNode, n)
	for i := 0; i < n; i++ {
		node := &tempNode{}
		fmt.Scanf("%d %d %d", &node.Addr, &node.Val, &node.Next)
		nodes[node.Addr] = node
	}

	headTempNode := nodes[head]
	headNode := &ListNode{Val: headTempNode.Val}
	curNode := headNode
	next := headTempNode.Next
	for next != -1 {
		node := nodes[next]
		curNode.Next = &ListNode{Val: node.Val}
		curNode = curNode.Next
		next = node.Next
	}

	fmt.Println(middleNode(headNode))
}
