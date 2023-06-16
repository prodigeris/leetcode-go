package _1_merge_two_sorted_lists

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var left, right int
	var result, last *ListNode
	for true {
		if list1 == nil {
			if last == nil {
				last = list2
				result = last
			} else {
				last.Next = list2
			}
			break
		}

		if list2 == nil {
			if last == nil {
				last = list1
				result = last
			} else {
				last.Next = list1
			}
			break
		}
		left = list1.Val
		right = list2.Val
		if left < right {
			if last != nil {
				last.Next = list1
			}
			last = list1
			list1 = list1.Next
		} else {
			if last != nil {
				last.Next = list2
			}
			last = list2
			list2 = list2.Next
		}
		if result == nil {
			result = last
		}
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}
