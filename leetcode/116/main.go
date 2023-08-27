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
	for stackLen := len(stack); stackLen != 0; stackLen = len(stack) {
		for i, node := range stack {
			if i+1 < stackLen {
				node.Next = stack[i+1]
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}

		if stackLen == len(stack) {
			break
		}

		stack = stack[stackLen:]
	}
	return root
}

func main() {
	connect(&Node{Val: 1, Left: &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}}, Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}}})
}
