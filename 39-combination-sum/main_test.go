package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected [][]int
	}{
		{
			name:     "Example 1",
			input:    []int{2, 3, 6, 7},
			target:   7,
			expected: [][]int{{2, 2, 3}, {7}},
		},
		{
			name:     "Example 2",
			input:    []int{2, 3, 5},
			target:   8,
			expected: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:     "Example 3",
			input:    []int{2},
			target:   1,
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := combinationSum(tt.input, tt.target)
			if !equivalentSetsOfSets(got, tt.expected) {
				t.Errorf("combinationSum(%v, %v) = %v; want %v", tt.input, tt.target, got, tt.expected)
			}
		})
	}
}

// Helper function to determine if two slices of slices have the same sets.
func equivalentSetsOfSets(a, b [][]int) bool {
	mapA, mapB := make(map[string]struct{}), make(map[string]struct{})

	for _, v := range a {
		mapA[toString(v)] = struct{}{}
	}
	for _, v := range b {
		mapB[toString(v)] = struct{}{}
	}

	return reflect.DeepEqual(mapA, mapB)
}

// Helper function to convert a slice of ints to a string representation.
func toString(a []int) string {
	return fmt.Sprint(a)
}
