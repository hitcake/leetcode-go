package tree

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	var root = &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	tree := BuildTree([]interface{}{1, 2, 3, nil, nil, 6, 7})
	if !IsSameTree(root, tree) {
		t.Errorf("IsSameTree(%v, %v) fail", tree, root)
	}
}

func TestPreOrder(t *testing.T) {
	var root = &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
	order := make([]int, 0)
	PreOrder(root, &order)
	if !reflect.DeepEqual(order, []int{1, 2, 3}) {
		t.Errorf("order = %v; want %v", order, []int{1, 2, 3})
	}
}

func TestPostOrder(t *testing.T) {
	var root = &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
	order := make([]int, 0)
	PostOrder(root, &order)
	if !reflect.DeepEqual(order, []int{2, 3, 1}) {
		t.Errorf("order = %v; want %v", order, []int{2, 3, 1})
	}
}

func TestInOrder(t *testing.T) {
	var root = &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
	order := make([]int, 0)
	InOrder(root, &order)
	if !reflect.DeepEqual(order, []int{2, 1, 3}) {
		t.Errorf("order = %v; want %v", order, []int{2, 1, 3})
	}
}
