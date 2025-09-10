package leetcode0102

import . "leetcode-go/tree"

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		newStack := make([]*TreeNode, 0)
		var list []int
		for _, node := range stack {
			list = append(list, node.Val)
			if node.Left != nil {
				newStack = append(newStack, node.Left)
			}
			if node.Right != nil {
				newStack = append(newStack, node.Right)
			}
		}
		result = append(result, list)
		stack = newStack
	}
	return result
}
