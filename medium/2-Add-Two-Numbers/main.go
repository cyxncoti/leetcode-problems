package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Intuition
// TC: O(max(n1, n2)), n1 and n2 are length of l1 and l2 respectively
// SC: O(1)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head, carry := l1, false
	for ; ; l1, l2 = l1.Next, l2.Next {
		l1.Val += l2.Val
		if carry {
			l1.Val++
		}
		carry = l1.Val > 9
		if carry {
			l1.Val -= 10
		}
		if l1.Next == nil {
			l1.Next = l2.Next
			break
		}
		if l2.Next == nil {
			break
		}
	}
	for carry {
		if l1.Next != nil {
			l1 = l1.Next
			if l1.Val == 9 {
				l1.Val = 0
			} else {
				l1.Val++
				break
			}
		} else {
			l1.Next = &ListNode{Val: 1, Next: nil}
			break
		}
	}
	return head
}
