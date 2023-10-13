package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return goDeeper(root, 0)
}

func goDeeper(root *TreeNode, level int) int {
	if root == nil {
		return level
	}

	level++
	l := goDeeper(root.Left, level)
	r := goDeeper(root.Right, level)
	if l > r {
		return l
	}
	return r
}
