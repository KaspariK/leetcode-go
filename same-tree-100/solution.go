package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func main() {
	a := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: nil,
			Right: nil,
		},
		Right: nil,
	}
	b := TreeNode{
			Val: 1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
				Left: nil,
				Right: nil,
			},
		}

	fmt.Println(IsSameTree(&a, &b))
}
