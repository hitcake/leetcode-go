package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListNode(nums []int) *ListNode {
	var head, p *ListNode
	for _, n := range nums {
		node := &ListNode{Val: n}
		if head == nil {
			head = node
			p = node
		} else {
			p.Next = node
			p = node
		}
	}
	return head
}

func DisplayListNode(head *ListNode) []int {
	nums := make([]int, 0)
	p := head
	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}
	return nums
}
