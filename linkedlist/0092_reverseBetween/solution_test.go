package leetcode0092

import (
	. "leetcode-go/linkedlist"
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	var tests = []struct {
		nums  []int
		left  int
		right int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{[]int{5}, 1, 1, []int{5}},
	}
	for _, test := range tests {
		head := BuildListNode(test.nums)
		got := DisplayListNode(reverseBetween(head, test.left, test.right))
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("expected: %v, got: %v", test.want, got)
		}
	}
}
