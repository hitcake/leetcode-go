package leetcode0083

import . "leetcode-go/linkedlist"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre, cur *ListNode = head, nil
	if pre.Next != nil {
		cur = pre.Next
	}
	for cur != nil {
		if pre.Val == cur.Val {
			pre.Next = cur.Next
			cur.Next = nil
			cur = pre.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return head
}
