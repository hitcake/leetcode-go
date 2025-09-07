package leetcode0061

import . "leetcode-go/linkedlist"

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	count := 0
	var tail *ListNode
	for p := head; p != nil; p = p.Next {
		count++
		if p.Next == nil {
			tail = p
		}
	}
	k %= count
	if k == 0 {
		return head
	}
	p := head
	for i := 1; i < count-k; i++ {
		p = p.Next
	}
	newHead := p.Next
	p.Next = nil
	tail.Next = head
	return newHead
}
