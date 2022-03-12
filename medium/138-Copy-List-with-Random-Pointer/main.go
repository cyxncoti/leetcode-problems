package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// Intuition: using a map
// TC: O(n), n = length of list
// SC: O(n), n = length of list
func firstSolution(head *Node) *Node {
	mapToNewNode := map[*Node]*Node{}
	sentinel := &Node{Val: 0, Next: head, Random: nil}
	for pred, current := sentinel, head; current != nil; current = current.Next {
		pred.Next = &Node{Val: current.Val, Next: nil, Random: nil}
		pred = pred.Next
		mapToNewNode[current] = pred
	}
	for n, o := sentinel.Next, head; o != nil; n, o = n.Next, o.Next {
		if o.Random != nil {
			n.Random = mapToNewNode[o.Random]
		}
	}
	return sentinel.Next
}

// Alternate duplicated node: without a map
// TC: O(n), n = length of list
// SC: O(n) but without a map, n = length of list
func secondSolution(head *Node) *Node {
	for current := head; current != nil; {
		next := current.Next
		current.Next = &Node{Val: current.Val, Next: next, Random: nil}
		current = next
	}
	for current := head; current != nil; current = current.Next.Next {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
	}
	sentinel := &Node{}
	for pred, current := sentinel, head; current != nil; current = current.Next {
		pred.Next = current.Next
		pred = pred.Next
		current.Next = current.Next.Next
	}
	return sentinel.Next
}
