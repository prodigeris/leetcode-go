package _1_merge_two_sorted_lists

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var curr, prev, start *ListNode
	for list1 != nil || list2 != nil {
		if list2 == nil || (list1 != nil && list1.Val <= list2.Val) {
			curr, list1 = list1, list1.Next
		} else {
			curr, list2 = list2, list2.Next
		}
		if prev != nil {
			prev.Next = curr
		}
		if start == nil {
			start = curr
		}
		prev = curr
	}
	return start
}

type ListNode struct {
	Val  int
	Next *ListNode
}
