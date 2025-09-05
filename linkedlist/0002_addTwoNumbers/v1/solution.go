package leetcode0002v1

import . "leetcode-go/linkedlist"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else {
		p := l1
		q := l2
		var head, tail *ListNode
		carry := 0
		for p != nil && q != nil {
			value := p.Val + q.Val + carry
			if value >= 10 {
				carry = 1
				value = value - 10
			} else {
				carry = 0
			}
			node := new(ListNode)
			node.Val = value
			if head == nil {
				head = node
				tail = node
			} else {
				tail.Next = node
				tail = node
			}
			p = p.Next
			q = q.Next
		}
		if p != nil {
			tail.Next = p
			tail = p
		} else if q != nil {
			tail.Next = q
			tail = q
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
		return head
	}

}
