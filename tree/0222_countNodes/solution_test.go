package leetcode0222

import (
	. "leetcode-go/tree"
	"reflect"
	"testing"
)

func TestCountNodes(t *testing.T) {
	var testCases = []struct {
		nums []interface{}
		want int
	}{
		{nums: []interface{}{1, 2, 3, 4}, want: 4},
		{nums: []interface{}{1, 2, 3, 4, 5, 6}, want: 6},
	}
	for _, testCase := range testCases {
		root := BuildTree(testCase.nums)
		got := countNodes(root)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("expected:%v, got:%v", testCase.want, got)
		}
	}
}
