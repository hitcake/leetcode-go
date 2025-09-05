package leetcode0148

import (
	. "leetcode-go/linkedlist"
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{[]int{5, 4, 3, 1, 2}, []int{1, 2, 3, 4, 5}},
		{[]int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},
	}
	for _, tt := range tests {
		head := BuildListNode(tt.nums)
		got := DisplayListNode(sortList(head))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v want %v", got, tt.want)
		}

	}
}
