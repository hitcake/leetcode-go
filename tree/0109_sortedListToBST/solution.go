package leetcode0109

import (
	. "leetcode-go/linkedlist"
	. "leetcode-go/tree"
)

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	return sortedListToBST1(head, nil)
}

func sortedListToBST1(head, tail *ListNode) *TreeNode {
	middle := findMiddle(head, tail)
	root := &TreeNode{Val: middle.Val}
	if head != middle {
		root.Left = sortedListToBST1(head, middle)
	}
	if middle.Next != tail {
		root.Right = sortedListToBST1(middle.Next, tail)
	}

	return root
}

func findMiddle(head, tail *ListNode) *ListNode {
	fast, slow := head, head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
