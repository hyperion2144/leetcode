package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	
	stack := []*Node{root}
	for len(stack) > 0 {
		copyStack := stack
		stack = nil

		for i := 0; i < len(copyStack); i++ {
			node := copyStack[i]
			if i + 1 < len(copyStack) {
				node.Next = copyStack[i+1]
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	return root
}
