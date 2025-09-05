package leetcode0002v2

import . "leetcode-go/linkedlist"

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	add(l1, l2, tail, 0)
	return head.Next
}

func add(l1, l2, tail *ListNode, carry int) {
	if l1 == nil || l2 == nil {
		if l1 != nil {
			tail.Next = l1
			tail = tail.Next
		} else if l2 != nil {
			tail.Next = l2
			tail = tail.Next
		} else if carry > 0 {
			tail.Next = &ListNode{}
			tail = tail.Next
		}
		for carry > 0 {
			tail.Val += carry
			if tail.Val >= 10 {
				carry = 1
				tail.Val = tail.Val - 10
			} else {
				carry = 0
			}
			if carry > 0 {
				if tail.Next != nil {
					tail = tail.Next
				} else {
					tail.Next = &ListNode{Val: 0}
					tail = tail.Next
				}
			}
		}
	} else {
		tail.Next = &ListNode{Val: l1.Val + l2.Val + carry}
		tail = tail.Next
		if tail.Val >= 10 {
			tail.Val -= 10
			carry = 1
		} else {
			carry = 0
		}
		add(l1.Next, l2.Next, tail, carry)
	}
}
