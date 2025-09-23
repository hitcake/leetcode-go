package leetcode0111

import (
	. "leetcode-go/tree"
	"testing"
)

func TestMinDepth(t *testing.T) {
	var tests = []struct {
		nums []interface{}
		want int
	}{
		{[]interface{}{1, 2, 3}, 2},
		{[]interface{}{1, nil, 3}, 1},
		{[]interface{}{3, 9, 20, nil, nil, 15, 7}, 2},
	}
	for _, tt := range tests {
		root := BuildTree(tt.nums)
		got := minDepth(root)
		if got != tt.want {
			t.Errorf("got: %v, want: %v", got, tt.want)
		}
	}
}
