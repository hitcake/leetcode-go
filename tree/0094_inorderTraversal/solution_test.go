package leetcode0094

import (
	. "leetcode-go/tree"
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want []int
	}{
		{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, []int{2, 1, 3}},
	}
	for _, tt := range tests {
		got := inorderTraversal(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}
