package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	*list = append(*list, root.Val)
	preOrder(root.Left, list)
	preOrder(root.Right, list)
}

func inOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, list)
	*list = append(*list, root.Val)
	inOrder(root.Right, list)
}

func postOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	postOrder(root.Left, list)
	postOrder(root.Right, list)
	*list = append(*list, root.Val)
}
