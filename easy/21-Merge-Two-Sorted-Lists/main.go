package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	for current := dummyHead; ; current = current.Next {
		if list1 == nil {
			current.Next = list2
			break
		}
		if list2 == nil {
			current.Next = list1
			break
		}
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
	}
	return dummyHead.Next
}
