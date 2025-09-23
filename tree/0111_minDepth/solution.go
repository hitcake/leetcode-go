package leetcode0111

import . "leetcode-go/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var leftDepth = minDepth(root.Left) + 1
	var rightDepth = minDepth(root.Right) + 1
	return min(leftDepth, rightDepth)
}
