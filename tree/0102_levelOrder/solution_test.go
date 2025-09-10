package leetcode0102

import (
	. "leetcode-go/tree"
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	var testCases = []struct {
		tree []interface{}
		want [][]int
	}{
		{tree: []interface{}{1, 2, 3, nil, nil, 5, 7}, want: [][]int{[]int{1}, []int{2, 3}, []int{5, 7}}},
		{tree: []interface{}{1}, want: [][]int{[]int{1}}},
		{tree: []interface{}{}, want: [][]int{}},
		{tree: []interface{}{1, 2, 3, 4, 5}, want: [][]int{[]int{1}, []int{2, 3}, []int{4, 5}}},
	}
	for _, testCase := range testCases {
		root := BuildTree(testCase.tree)
		got := levelOrder(root)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("expected:%v, got:%v", testCase.want, got)
		}
	}
}
