package leetcode0086

import (
	. "leetcode-go/linkedlist"
	v2 "leetcode-go/linkedlist/0086_partition/v2"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	var tests = []struct {
		nums []int
		x    int
		want []int
	}{
		{[]int{1, 5, 3, 2, 7}, 3, []int{1, 2, 5, 3, 7}},
		{[]int{1, 4, 3, 2, 5, 2}, 3, []int{1, 2, 2, 4, 3, 5}},
		{[]int{2, 1}, 2, []int{1, 2}},
	}
	//for _, tt := range tests {
	//	head := BuildListNode(tt.nums)
	//	got := DisplayListNode(v1.Partition(head, tt.x))
	//	if !reflect.DeepEqual(got, tt.want) {
	//		t.Errorf("expected: %v, got: %v", tt.want, got)
	//	}
	//}
	for _, tt := range tests {
		head := BuildListNode(tt.nums)
		got := DisplayListNode(v2.Partition(head, tt.x))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("expected: %v, got: %v", tt.want, got)
		}
	}
}
