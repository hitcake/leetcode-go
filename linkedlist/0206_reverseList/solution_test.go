package leetcode0206

import (
	. "leetcode-go/linkedlist"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{1, 2}, []int{2, 1}},
	}
	for _, tt := range tests {
		head := BuildListNode(tt.nums)
		newHead := reverseList(head)
		got := DisplayListNode(newHead)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("expected: %v, got: %v", tt.want, got)
		}
	}
}
