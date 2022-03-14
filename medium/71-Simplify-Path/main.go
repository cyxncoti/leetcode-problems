package main

import "strings"

func simplifyPath(path string) string {
	type ListNode struct {
		Val  string
		Prev *ListNode
		Next *ListNode
	}
	sentinel := &ListNode{}
	stack := sentinel
	splits := strings.Split(path, "/")
	for _, s := range splits {
		switch s {
		case "", ".":
		case "..":
			if stack != sentinel {
				stack = stack.Prev
			}
		default:
			stack.Next = &ListNode{Val: s, Prev: stack, Next: nil}
			stack = stack.Next
		}
	}
	if stack == sentinel {
		return "/"
	}
	stack.Next = nil
	var sb strings.Builder
	for stack = sentinel.Next; stack != nil; stack = stack.Next {
		sb.WriteRune('/')
		sb.WriteString(stack.Val)
	}
	return sb.String()
}
