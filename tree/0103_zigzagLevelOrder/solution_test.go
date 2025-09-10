package leetcode0103

import (
	. "leetcode-go/tree"
	"reflect"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	var tests = []struct {
		nums []interface{}
		want [][]int
	}{
		{nums: []interface{}{1, 2, 3, nil, nil, 6, 7}, want: [][]int{{1}, {3, 2}, {6, 7}}},
		{nums: []interface{}{1}, want: [][]int{{1}}},
		{nums: []interface{}{}, want: [][]int{}},
	}
	for _, tt := range tests {
		root := BuildTree(tt.nums)
		got := zigzagLevelOrder(root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
		}
	}
}
