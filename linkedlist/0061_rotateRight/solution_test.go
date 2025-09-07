package leetcode0061

import (
	. "leetcode-go/linkedlist"
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, []int{5, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{0, 1, 2}, 4, []int{2, 0, 1}},
	}

	for _, tt := range tests {
		head := BuildListNode(tt.nums)
		got := DisplayListNode(rotateRight(head, tt.k))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("input: %v, got: %v, want: %v", tt.nums, got, tt.want)
		}
	}
}
