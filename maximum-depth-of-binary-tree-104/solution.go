package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	return maxInt(left, right) + 1
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}