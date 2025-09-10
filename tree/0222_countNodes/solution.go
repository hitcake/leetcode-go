package leetcode0222

import (
	. "leetcode-go/tree"
	"math"
)

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else {
		leftHeight := 0
		for p := root.Left; p != nil; p = p.Left {
			leftHeight++
		}
		rightHeight := 0
		for p := root.Right; p != nil; p = p.Right {
			rightHeight++
		}
		if leftHeight == rightHeight {
			return int(math.Pow(2, float64(leftHeight+1))) - 1
		} else {
			return countNodes(root.Left) + countNodes(root.Right) + 1
		}
	}
}
