package removeNthFromEnd

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	x := 0
	for p != nil {
		x++
		p = p.Next
	}

	m := x - n
	if m <= 0 {
		return head.Next
	}
	p = head
	for i := 0; i < m-1; i++ {
		p = p.Next
	}
	q := p.Next
	if q != nil {
		p.Next = q.Next
	}
	return head
}
