package leetcode0109

import (
	. "leetcode-go/linkedlist"
	. "leetcode-go/tree"
	"testing"
)

func TestSortListToBST(t *testing.T) {
	var listNode = BuildListNode([]int{1, 2, 3, 4, 5, 6, 7})
	got := sortedListToBST(listNode)
	want := BuildTree([]interface{}{4, 2, 6, 1, 3, 5, 7})
	if !IsSameTree(got, want) {
		t.Error("got 1  ")
	}
	listNode = BuildListNode([]int{-10, -3, 0, 5, 9})
	got = sortedListToBST(listNode)
	want = BuildTree([]interface{}{0, -3, 9, -10, nil, 5})
	if !IsSameTree(got, want) {
		t.Error("got 2  ")
	}
}
