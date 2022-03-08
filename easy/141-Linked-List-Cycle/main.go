package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	return secondSolution(head)
}

// Brute Force
// TC: O(n), n = number of nodes in the list
// SC: O(n), n = number of nodes in the list
func firstSolution(head *ListNode) bool {
	for seen := map[*ListNode]bool{}; head != nil; head = head.Next {
		if seen[head] {
			return true
		}
		seen[head] = true
	}
	return false
}

// Floyd's Tortoise and Hare
// TC: O(n), n = number of nodes in the list
// SC: O(1)
func secondSolution(head *ListNode) bool {
	for hare := head; hare != nil && hare.Next != nil; {
		hare = hare.Next.Next
		head = head.Next
		if hare == head {
			return true
		}
	}
	return false
}
