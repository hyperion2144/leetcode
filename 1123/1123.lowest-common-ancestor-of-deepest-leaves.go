package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1123 lang=golang
 * @lcpr version=21913
 *
 * [1123] 最深叶节点的最近公共祖先
 *
 * https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/description/
 *
 * algorithms
 * Medium (75.18%)
 * Likes:    203
 * Dislikes: 0
 * Total Accepted:    21.4K
 * Total Submissions: 28.5K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]'
 *
 * 给你一个有根节点 root 的二叉树，返回它 最深的叶节点的最近公共祖先 。
 *
 * 回想一下：
 *
 *
 * 叶节点 是二叉树中没有子节点的节点
 * 树的根节点的 深度 为 0，如果某一节点的深度为 d，那它的子节点的深度就是 d+1
 * 如果我们假定 A 是一组节点 S 的 最近公共祖先，S 中的每个节点都在以 A 为根节点的子树中，且 A 的深度达到此条件下可能的最大值。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：root = [3,5,1,6,2,0,8,null,null,7,4]
 * 输出：[2,7,4]
 * 解释：我们返回值为 2 的节点，在图中用黄色标记。
 * 在图中用蓝色标记的是树的最深的节点。
 * 注意，节点 6、0 和 8 也是叶节点，但是它们的深度是 2 ，而节点 7 和 4 的深度是 3 。
 *
 *
 * 示例 2：
 *
 * 输入：root = [1]
 * 输出：[1]
 * 解释：根节点是树中最深的节点，它是它本身的最近公共祖先。
 *
 *
 * 示例 3：
 *
 * 输入：root = [0,1,3,null,2]
 * 输出：[2]
 * 解释：树中最深的叶节点是 2 ，最近公共祖先是它自己。
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数将在 [1, 1000] 的范围内。
 * 0 <= Node.val <= 1000
 * 每个节点的值都是 独一无二 的。
 *
 *
 *
 *
 * 注意：本题与力扣 865
 * 重复：https://leetcode-cn.com/problems/smallest-subtree-with-all-the-deepest-nodes/
 *
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, node := dfs(root)
	return node
}

func dfs(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	left, leftNode := dfs(root.Left)
	right, rightNode := dfs(root.Right)

	if left > right {
		return left + 1, leftNode
	}

	if right > left {
		return right + 1, rightNode
	}

	return left + 1, root
}

// @lc code=end

/*
// @lcpr case=start
// [3,5,1,6,2,0,8,null,null,7,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,3,null,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,5,6,3,21,20,7,49,4,14,244,null,181,47,9,10,62,155,17,8,60,null,null,287,194,182,null,50,30,35,19,11,82,73,164,172,108,25,12,24,156,77,null,null,203,null,null,202,74,53,42,48,232,36,137,66,27,13,84,90,135,null,237,236,184,255,186,284,40,39,88,26,280,54,null,null,91,187,225,209,null,null,94,113,59,65,138,87,67,null,243,null,240,38,196,null,167,68,28,70,23,15,98,101,131,139,212,175,null,null,null,null,238,214,267,null,248,null,null,null,177,46,55,41,107,176,44,51,null,null,205,null,null,125,null,242,null,291,null,293,126,null,274,199,257,71,97,102,null,null,null,null,null,83,null,null,null,null,null,85,null,null,261,null,96,140,32,61,104,79,57,29,16,18,198,173,null,null,null,null,null,154,null,null,null,null,null,276,251,null,279,null,null,null,289,200,92,null,58,260,56,78,221,208,null,null,109,null,99,null,null,null,null,null,null,null,292,null,null,null,153,null,null,282,null,null,null,null,null,297,null,null,null,233,145,null,144,null,288,null,119,null,162,null,null,null,75,105,null,170,95,210,103,216,null,31,22,null,37,43,null,null,null,190,null,null,null,null,null,null,null,null,296,null,null,265,150,null,null,81,null,null,143,266,207,192,null,null,226,null,115,118,null,null,null,null,null,163,null,null,null,null,null,null,160,null,null,178,null,null,null,159,null,null,null,null,127,null,253,191,180,null,null,null,null,132,null,null,null,null,34,33,89,52,111,45,null,262,null,null,null,null,null,211,null,193,null,183,null,null,null,null,null,null,null,294,271,null,null,227,283,270,179,null,189,222,null,null,215,166,null,null,null,null,null,null,281,134,188,63,148,136,129,168,null,80,114,null,69,64,null,null,223,null,277,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,201,259,null,258,null,null,null,null,null,null,null,null,null,72,264,218,null,null,158,133,142,268,null,93,128,123,161,230,76,86,106,295,null,null,null,249,null,null,null,null,263,235,null,null,null,null,null,null,null,null,null,null,null,null,null,121,130,null,152,null,124,219,247,null,234,239,100,110,117,116,151,298,null,null,null,null,null,null,null,217,122,185,null,null,null,null,204,null,null,null,290,null,null,null,null,null,112,197,269,120,195,null,147,157,206,null,null,null,null,null,141,245,null,null,null,null,null,220,165,null,null,null,null,null,229,null,null,228,231,null,213,null,null,169,146,null,300,224,null,null,null,null,273,null,286,246,null,null,252,null,171,149,null,null,null,241,278,null,null,null,null,299,null,null,null,null,null,174,275,null,null,null,null,null,null,null,250,null,null,256,254,null,272,null,null,null,285]\n
// @lcpr case=end

*/

func main() {
	fmt.Println(lcaDeepestLeaves(&TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}}))
}
