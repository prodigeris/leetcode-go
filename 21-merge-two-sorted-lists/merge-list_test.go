package _1_merge_two_sorted_lists

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l  *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example 1", args{sliceToLinkNode([]int{1, 2, 4}), sliceToLinkNode([]int{1, 3, 4})},
			sliceToLinkNode([]int{1, 1, 2, 3, 4, 4}),
		},
		{"Example 2", args{sliceToLinkNode([]int{}), sliceToLinkNode([]int{})},
			sliceToLinkNode([]int{}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := mergeTwoLists(tt.args.l, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists got %v, wanted %v", listNodeToSlice(got), listNodeToSlice(tt.want))

			}
		})
	}
}

func listNodeToSlice(node *ListNode) []int {
	s := make([]int, 0)
	for node != nil {
		s = append(s, node.Val)
		node = node.Next
	}
	return s
}

func sliceToLinkNode(slice []int) *ListNode {
	var head *ListNode
	for i := len(slice) - 1; i >= 0; i-- {
		head = &ListNode{slice[i], head}
	}
	return head
}
