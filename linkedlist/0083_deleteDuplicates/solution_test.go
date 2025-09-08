package leetcode0083

import (
	. "leetcode-go/linkedlist"
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 1}, []int{0, 1}},
		{[]int{0, 1, 2, 3, 3}, []int{0, 1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{1, 1, 1, 1}, []int{1}},
	}
	for _, test := range tests {
		head := BuildListNode(test.nums)
		got := DisplayListNode(deleteDuplicates(head))
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("expected:%v, got:%v", test.want, got)
		}
	}
}
