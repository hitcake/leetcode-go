package tree

import (
	"reflect"
	"testing"
)

func TestPreOrder(t *testing.T) {
	var root = &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
	order := make([]int, 0)
	preOrder(root, &order)
	if !reflect.DeepEqual(order, []int{1, 2, 3}) {
		t.Errorf("order = %v; want %v", order, []int{1, 2, 3})
	}
}

func TestPostOrder(t *testing.T) {
	var root = &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
	order := make([]int, 0)
	postOrder(root, &order)
	if !reflect.DeepEqual(order, []int{2, 3, 1}) {
		t.Errorf("order = %v; want %v", order, []int{2, 3, 1})
	}
}

func TestInOrder(t *testing.T) {
	var root = &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
	order := make([]int, 0)
	inOrder(root, &order)
	if !reflect.DeepEqual(order, []int{2, 1, 3}) {
		t.Errorf("order = %v; want %v", order, []int{2, 1, 3})
	}
}
