package leetcode0107

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
		{tree: []interface{}{1, 2, 3, nil, nil, 5, 7}, want: [][]int{{5, 7}, {2, 3}, {1}}},
		{tree: []interface{}{1}, want: [][]int{{1}}},
		{tree: []interface{}{}, want: [][]int{}},
		{tree: []interface{}{1, 2, 3, 4, 5}, want: [][]int{{4, 5}, {2, 3}, {1}}},
	}
	for _, testCase := range testCases {
		root := BuildTree(testCase.tree)
		got := levelOrderBottom(root)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("expected:%v, got:%v", testCase.want, got)
		}
	}
}
