package main

func isValid(s string) bool {
	type ListNode struct {
		Val  rune
		Next *ListNode
	}
	var stack *ListNode
	for _, r := range s {
		switch r {
		case '(':
			stack = &ListNode{Val: ')', Next: stack}
		case '{':
			stack = &ListNode{Val: '}', Next: stack}
		case '[':
			stack = &ListNode{Val: ']', Next: stack}
		case ')', '}', ']':
			if stack == nil || stack.Val != r {
				return false
			}
			stack = stack.Next
		}
	}
	if stack != nil {
		return false
	}
	return true
}
