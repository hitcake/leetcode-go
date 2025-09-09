package leetcode0206

import . "leetcode-go/linkedlist"

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, cur := head, head.Next
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		if pre == head {
			pre.Next = nil
		}
		pre = cur
		cur = tmp
	}
	return pre
}
