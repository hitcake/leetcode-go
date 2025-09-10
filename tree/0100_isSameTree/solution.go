package leetcode0100

import . "leetcode-go/tree"

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
