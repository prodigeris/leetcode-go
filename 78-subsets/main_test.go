package main

import (
	"fmt"
	"reflect"
	sort "sort"
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

		{
			name:  "Example 3",
			input: []int{9, 0, 3, 5, 7},
			output: [][]int{
				{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5},
				{5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5},
				{0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9},
				{0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7},
				{0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9},
			},
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := subsets(tt.input)
			if !equivalentSets(got, tt.output) {
				t.Errorf("subsets(%v) = %v; want %v", tt.input, got, tt.output)
			}
		})
	}
}

// Helper function to determine if two slices of slices have the same sets.
func equivalentSets(a, b [][]int) bool {
	mapA, mapB := make(map[string]struct{}), make(map[string]struct{})

	for _, v := range a {
		sort.Ints(v)
		mapA[toString(v)] = struct{}{}
	}
	for _, v := range b {
		sort.Ints(v)
		mapB[toString(v)] = struct{}{}
	}

	return reflect.DeepEqual(mapA, mapB)
}

// Helper function to convert a slice of ints to a string representation.
func toString(a []int) string {
	return fmt.Sprint(a)
}
