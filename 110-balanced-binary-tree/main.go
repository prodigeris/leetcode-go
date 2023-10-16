package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	stillBalanced := true
	dfs(root, &stillBalanced)
	return stillBalanced
}

func dfs(root *TreeNode, stillBalanced *bool) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left, stillBalanced)
	r := dfs(root.Right, stillBalanced)

	isBalanced := math.Abs(float64(l-r)) < 2
	if !isBalanced {
		*stillBalanced = false
	}

	return 1 + maximum(l, r)
}

func maximum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
