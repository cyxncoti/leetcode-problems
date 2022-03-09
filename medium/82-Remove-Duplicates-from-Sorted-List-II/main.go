package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Intuition
// TC: O(n), n = number of nodes in given list
// SC: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	sentinel := &ListNode{Val: 0, Next: head}
	pred := sentinel
	for head != nil {
		for head.Next != nil && head.Next.Val == head.Val {
			head = head.Next
		}
		if head != pred.Next {
			pred.Next = head.Next
		} else {
			pred = pred.Next
		}
		head = head.Next
	}
	return sentinel.Next
}
