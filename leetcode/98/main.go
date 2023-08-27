package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LINK - https://leetcode-cn.com/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, nil, nil)
}

func isValidBSTHelper(node *TreeNode, min *TreeNode, max *TreeNode) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Val <= min.Val) || (max != nil && node.Val >= max.Val) {
		return false
	}

	return isValidBSTHelper(node.Left, min, node) && isValidBSTHelper(node.Right, node, max)
}
