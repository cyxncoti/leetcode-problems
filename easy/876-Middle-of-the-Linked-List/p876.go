package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fast := head.Next
	for {
		if fast == nil {
			return head
		}
		head = head.Next
		fast = fast.Next
		if fast == nil {
			return head
		}
		fast = fast.Next
	}
}
