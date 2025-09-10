package leetcode0103

import . "leetcode-go/tree"

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for number := 0; len(stack) > 0; number++ {
		var level []int
		newStack := make([]*TreeNode, 0)
		for _, node := range stack {
			level = append(level, node.Val)
			if node.Left != nil {
				newStack = append(newStack, node.Left)
			}
			if node.Right != nil {
				newStack = append(newStack, node.Right)
			}
		}
		stack = newStack
		if number%2 == 1 {
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		}
		result = append(result, level)
	}
	return result
}
