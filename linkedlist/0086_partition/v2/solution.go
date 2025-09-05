package leetcode0086v2

import . "leetcode-go/linkedlist"

func Partition(head *ListNode, x int) *ListNode {
	pHead := &ListNode{}
	pTail := pHead

	qHead := &ListNode{}
	qTail := qHead

	for head != nil {
		if head.Val < x {
			pTail.Next = head
			pTail = pTail.Next
		} else {
			qTail.Next = head
			qTail = qTail.Next
		}
		head = head.Next
	}
	pTail.Next = qHead.Next
	qTail.Next = nil
	return pHead.Next
}
