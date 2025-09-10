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
		levelStack := stack[:]
		stack = make([]*TreeNode, 0)
		levelValues := make([]int, 0)
		for _, node := range levelStack {
			levelValues = append(levelValues, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		result = append(result, levelValues)
	}
	return result
}
