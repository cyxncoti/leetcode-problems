package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	maxWidth := 0
	for nodes := []*TreeNode{root}; len(nodes) > 0; {
		min, max, nextNodes := nodes[0].Val, nodes[len(nodes)-1].Val, []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				node.Left.Val = 2 * (node.Val - min)
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = 2*(node.Val-min) + 1
				nextNodes = append(nextNodes, node.Right)
			}
		}
		if width := max - min + 1; width > maxWidth {
			maxWidth = width
		}
		nodes = nextNodes
	}
	return maxWidth
}
