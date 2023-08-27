package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	tree := &TreeNode{
		Val: preorder[0],
	}

	for i, v := range inorder {
		if v == preorder[0] {
			tree.Left = buildTree(preorder[1:i+1], inorder[:i])
			tree.Right = buildTree(preorder[i+1:], inorder[i+1:])
			break
		}
	}

	return tree
}
