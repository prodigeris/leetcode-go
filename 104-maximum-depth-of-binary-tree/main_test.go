package main

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"#1", &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, 3},
		{"#2", &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, 2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := maxDepth(tt.root); got != tt.want {
				t.Errorf("%v failed; got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
