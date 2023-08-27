package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	slow, fast, distance := head, head, 0
	var preNode, nextNode *ListNode = nil, nil
	for fast != nil {
		fast = fast.Next
		distance++

		if distance == k-1 {
			nextNode = fast.Next
			fast.Next = nil

			var pre, cur *ListNode = nil, slow
			for cur != nil {
				next := cur.Next
				cur.Next = pre
				pre = cur
				cur = next
			}

			if preNode != nil {
				preNode.Next = fast
			} else {
				head = pre
			}

			slow.Next = nextNode

			preNode = slow
			distance = 0
			slow = nextNode
			fast = nextNode
		}

	}

	return head
}

func main() {
	reverseKGroup(&ListNode{Val: 1, Next: &ListNode{Val: 2, }}, 2)
}
