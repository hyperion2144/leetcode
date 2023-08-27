package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	nodeIndex := make(map[int]*ListNode)
	for i, p := 0, head; p != nil; i, p = i+1, p.Next {
		nodeIndex[i] = p
	}

	k = k % len(nodeIndex)
	if k == 0 {
		return head
	}

	nodeIndex[len(nodeIndex)-1].Next = head
	nodeIndex[len(nodeIndex)-k-1].Next = nil
	return nodeIndex[len(nodeIndex)-k]
}
