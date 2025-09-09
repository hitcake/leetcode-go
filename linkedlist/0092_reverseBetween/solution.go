package leetcode0092

import . "leetcode-go/linkedlist"

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	n := right - left + 1
	values := make([]int, n)
	index := 0
	for p := head; p != nil; p = p.Next {
		index++
		if index < left {
			continue
		} else if index >= left && index <= right {
			values[index-left] = p.Val
		} else {
			break
		}
	}
	index = 0
	for p := head; p != nil; p = p.Next {
		index++
		if index < left {
			continue
		} else if index >= left && index <= right {
			p.Val = values[n-(index-left)-1]
		} else {
			break
		}
	}
	return head

}
