package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(float64(
		dfs(root.Left)-
			dfs(root.Right))) < 2
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + maximum(dfs(root.Left), dfs(root.Right))
}

func maximum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
