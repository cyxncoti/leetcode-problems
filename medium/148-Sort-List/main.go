package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := createList([]int{-1, 5, 3, 4, 0})
	list = sortList(list)
	printList(list, "Success")
}

func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{
		Val:  vals[0],
		Next: nil,
	}
	for i, n := 1, head; i < len(vals); i++ {
		n.Next = &ListNode{
			Val:  vals[i],
			Next: nil,
		}
		n = n.Next
	}
	return head
}

func printList(n *ListNode, msg string) {
	fmt.Printf(msg + ":")
	for n != nil {
		fmt.Printf(" %d,", n.Val)
		n = n.Next
	}
	fmt.Println()
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Val: 0, Next: head}
	for sublistSize := 1; ; sublistSize *= 2 {
		p1, p2 := dummyHead.Next, getNextNNode(dummyHead.Next, sublistSize)
		if p2 == nil {
			return dummyHead.Next
		}
		for tail := dummyHead; p2 != nil; {
			head = getNextNNode(p2, sublistSize) // Get the head of next sublist beforehand
			tail = mergeSublists(tail, p1, p2, sublistSize)
			p1, p2 = head, getNextNNode(head, sublistSize)
		}
	}
}

func getNextNNode(cur *ListNode, n int) *ListNode {
	for i := 0; i < n; i++ {
		if cur == nil {
			return nil
		}
		cur = cur.Next
	}
	return cur
}

func mergeSublists(tail *ListNode, p1 *ListNode, p2 *ListNode, sublistSize int) *ListNode {
	for c1, c2 := 0, 0; ; {
		if c1 == sublistSize || p1 == nil {
			tail.Next = p2
			return getNextNNode(tail, sublistSize-c2)
		}
		if c2 == sublistSize || p2 == nil {
			tail.Next = p1
			tail = getNextNNode(tail, sublistSize-c1)
			tail.Next = p2 // Prevent from creating cycle
			return tail
		}
		if p1.Val < p2.Val {
			tail.Next = p1
			p1 = p1.Next
			c1++
		} else {
			tail.Next = p2
			p2 = p2.Next
			c2++
		}
		tail = tail.Next
	}
}
