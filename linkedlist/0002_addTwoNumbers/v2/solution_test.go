package leetcode0002v2

import (
	"fmt"
	. "leetcode-go/linkedlist"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	node := addTwoNumbers(int2ListNode(9999999), int2ListNode(9999))
	fmt.Println(listNode2Int(node))
}

func int2ListNode(value int) *ListNode {
	var head, tail *ListNode
	for value != 0 {
		val := value % 10
		value = value / 10
		p := &ListNode{Val: val}
		if head == nil {
			head = p
			tail = p
		} else {
			tail.Next = p
			tail = tail.Next
		}
	}
	return head
}
func listNode2Int(head *ListNode) int {
	p := head
	i := 1
	result := 0
	for p != nil {
		result += p.Val * i
		p = p.Next
		i *= 10
	}
	return result
}
