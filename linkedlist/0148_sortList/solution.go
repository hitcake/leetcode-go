package leetcode0148

import . "leetcode-go/linkedlist"

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	for p := head; p != nil; p = p.Next {
		for q := p.Next; q != nil; q = q.Next {
			if q.Val < p.Val {
				p.Val, q.Val = q.Val, p.Val
			}
		}
	}
	return head
}
