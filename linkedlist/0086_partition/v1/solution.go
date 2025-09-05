package leetcode0086v1

import . "leetcode-go/linkedlist"

func Partition(head *ListNode, x int) *ListNode {
	var newhead, tail *ListNode
	for p := head; p != nil; p = p.Next {
		if p.Val < x {
			node := &ListNode{Val: p.Val}
			if tail == nil {
				newhead = node
				tail = node
			} else {
				tail.Next = node
				tail = node
			}
		}
	}
	for p := head; p != nil; p = p.Next {
		if p.Val >= x {
			node := &ListNode{Val: p.Val}
			if tail == nil {
				newhead = node
				tail = node
			} else {
				tail.Next = node
				tail = node
			}
		}
	}
	return newhead
}
