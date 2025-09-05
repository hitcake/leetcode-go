package leetcode0019

import (
	. "leetcode-go/linkedlist"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	var tests = []struct {
		nums []int
		n    int
		want []int
	}{
		{nums: []int{1}, n: 1, want: []int{}},
		{nums: []int{1, 2, 3, 4, 5}, n: 2, want: []int{1, 2, 3, 5}},
		{nums: []int{1, 2}, n: 1, want: []int{1}},
		{nums: []int{1, 2, 3}, n: 3, want: []int{2, 3}},
	}
	for _, test := range tests {
		got := DisplayListNode(removeNthFromEnd(BuildListNode(test.nums), test.n))
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("expected: %v, got: %v", test.want, got)
		}
	}
}
