package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(nums []interface{}) *TreeNode {
	return buildTreeChild(nums, 0)
}

func IsSameTree(p, q *TreeNode) bool {
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

func buildTreeChild(nums []interface{}, index int) *TreeNode {
	if index >= len(nums) {
		return nil
	}
	if nums[index] == nil {
		return nil
	}
	var root *TreeNode
	switch t := nums[index].(type) {
	case int:
		root = &TreeNode{Val: t}
	default:
		fmt.Printf("nums[%d]=%v", index, nums[index])
		return nil
	}

	root.Left = buildTreeChild(nums, 2*index+1)
	root.Right = buildTreeChild(nums, 2*index+2)
	return root
}

func PreOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	*list = append(*list, root.Val)
	PreOrder(root.Left, list)
	PreOrder(root.Right, list)
}

func InOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	InOrder(root.Left, list)
	*list = append(*list, root.Val)
	InOrder(root.Right, list)
}

func PostOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	PostOrder(root.Left, list)
	PostOrder(root.Right, list)
	*list = append(*list, root.Val)
}
