package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return dfs(root) != -1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left)
	if l == -1 {
		return -1
	}
	r := dfs(root.Right)
	if r == -1 {
		return -1
	}

	if math.Abs(float64(l-r)) < 2 {
		return -1
	}

	return 1 + maximum(l, r)
}

func maximum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
