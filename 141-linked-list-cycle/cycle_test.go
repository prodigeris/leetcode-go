package main

import (
	"testing"
)

func TestHasCycle(t *testing.T) {

	pos := &ListNode{Val: 2}
	head1 := &ListNode{Val: 3, Next: pos}
	pos.Next = &ListNode{Val: 0, Next: &ListNode{Val: 4, Next: pos}}

	tests := []struct {
		head *ListNode
		want bool
	}{
		{head1, true},
	}

	for _, tt := range tests {
		got := hasCycle(tt.head)
		if got != tt.want {
			t.Errorf("HasCycle = %v; want %v", got, tt.want)
		}
	}
}

func TestHasCycleTortoise(t *testing.T) {

	pos := &ListNode{Val: 2}
	head1 := &ListNode{Val: 3, Next: pos}
	pos.Next = &ListNode{Val: 0, Next: &ListNode{Val: 4, Next: pos}}

	tests := []struct {
		head *ListNode
		want bool
	}{
		{head1, true},
	}

	for _, tt := range tests {
		got := hasCycleTortoise(tt.head)
		if got != tt.want {
			t.Errorf("HasCycle = %v; want %v", got, tt.want)
		}
	}
}
