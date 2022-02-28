package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	n := head
	for {
		if list1 == nil {
			n.Next = list2
			break
		}
		if list2 == nil {
			n.Next = list1
			break
		}
		if list1.Val < list2.Val {
			n.Next = list1
			list1 = list1.Next
		} else {
			n.Next = list2
			list2 = list2.Next
		}
		n = n.Next
	}
	return head.Next
}
