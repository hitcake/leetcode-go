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
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, want: []int{2, 1, 3}},
	}
	for _, tt := range tests {
		got := inorderTraversal(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}
