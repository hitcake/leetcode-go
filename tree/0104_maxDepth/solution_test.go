package leetcode0104

import (
	. "leetcode-go/tree"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want int
	}{
		{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, want: 2},
	}
	for _, tt := range tests {
		got := maxDepth(tt.root)
		if got != tt.want {
			t.Errorf("got: %v, want: %v", got, tt.want)
		}
	}
}
