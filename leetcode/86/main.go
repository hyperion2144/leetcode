package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := &ListNode{Next: head}
	cur := newHead
	pre := newHead
	next := newHead.Next

	for next != nil {
		if next.Val < x {
			if cur.Next == next {
				cur = cur.Next
				pre = pre.Next
				next = next.Next
				continue
			}
			temp := cur.Next
			pre.Next = next.Next
			cur.Next = next
			next.Next = temp
			cur = cur.Next
			next = pre.Next
		} else {
			pre = pre.Next
			next = next.Next
		}
	}

	return newHead.Next
}

func main() {
	partition(&ListNode{1, &ListNode{4, &ListNode{3, &ListNode{0, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}}, 3)
}
