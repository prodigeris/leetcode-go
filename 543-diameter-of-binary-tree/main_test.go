package main

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: 3,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		got := diameterOfBinaryTree(tt.root)
		if got != tt.want {
			t.Errorf("diameterOfBinaryTree(%v) = %v, want %v", tt.root, got, tt.want)
		}
	}
}
