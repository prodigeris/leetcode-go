package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	m := 0
	dfs(root, &m)
	return m
}

func dfs(root *TreeNode, m *int) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left, m)
	r := dfs(root.Right, m)
	*m = max(*m, l+r)

	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
