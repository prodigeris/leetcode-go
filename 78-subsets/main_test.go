package main

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output [][]int
	}{
		{
			name:   "Example 1",
			input:  []int{1, 2, 3},
			output: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:   "Example 2",
			input:  []int{0},
			output: [][]int{{}, {0}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := subsets(tt.input)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("subsets(%v) = %v; want %v", tt.input, got, tt.output)
			}
		})
	}
}
