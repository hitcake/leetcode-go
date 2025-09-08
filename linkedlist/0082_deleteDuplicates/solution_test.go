package leetcode082

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
		{[]int{1, 1, 1}, []int{}},
		{[]int{1, 1, 2}, []int{2}},
		{[]int{1, 2, 3, 3, 4, 4, 5}, []int{1, 2, 5}},
		{[]int{1, 1, 1, 2, 5}, []int{2, 5}},
	}
	for _, c := range tests {
		head := BuildListNode(c.nums)
		got := DisplayListNode(deleteDuplicates(head))
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("expected:%v, got:%v", c.want, got)
		}
	}
}
