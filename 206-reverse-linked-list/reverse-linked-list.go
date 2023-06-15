package _06_reverse_linked_list

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var p *ListNode
	for head != nil {
		head.Next, p, head = p, head, head.Next
	}

	return p
}

type ListNode struct {
	Val  int
	Next *ListNode
}
