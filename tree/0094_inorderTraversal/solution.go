package leetcode0094

import . "leetcode-go/tree"

func inorderTraversal(root *TreeNode) []int {
	var inorder func(node *TreeNode)
	var res []int
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}
