package main

func validateStackSequences(pushed []int, popped []int) bool {
	type ListNode struct {
		Val  int
		Next *ListNode
	}
	var stack *ListNode
	pop := 0
	for _, push := range pushed {
		stack = &ListNode{Val: push, Next: stack}
		for stack != nil && stack.Val == popped[pop] {
			stack = stack.Next
			pop++
		}
	}
	if stack != nil {
		return false
	}
	return true
}

func secondSolution(pushed []int, popped []int) bool {
	stack := -1
	for push, pop, n := 0, 0, len(pushed); push < n; push++ {
		stack++
		pushed[stack] = pushed[push]
		for stack >= 0 && pushed[stack] == popped[pop] {
			stack--
			pop++
		}
	}
	if stack >= 0 {
		return false
	}
	return true
}
