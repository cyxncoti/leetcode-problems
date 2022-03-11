package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	return secondSolution(head, k)
}

// Slow & Fast Pointer
// TC: O(n)
// SC: O(1)
func firstSolution(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast, n := head, head, 1
	for fast.Next != nil && n <= k {
		fast = fast.Next
		n++
	}
	if n > k {
		for fast.Next != nil {
			fast = fast.Next
			slow = slow.Next
		}
	} else {
		k %= n
		fast = head
		for i := 0; i < k; i++ {
			fast = fast.Next
		}
		for fast.Next != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}
	fast.Next = head
	head = slow.Next
	slow.Next = nil
	return head
}

// Count length and Cycle the List
// TC: O(n)
// SC: O(1)
func secondSolution(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	tail, n := head, 1
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	tail.Next = head
	k %= n
	for i := 0; i < n-k; i++ {
		tail = tail.Next
	}
	head = tail.Next
	tail.Next = nil
	return head
}
