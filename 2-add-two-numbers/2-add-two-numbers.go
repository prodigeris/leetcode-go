package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s := ListNode{}
	t := &s
	carry := 0
	for {
		if l1 == nil {
			l1 = &ListNode{}
		}
		if l2 == nil {
			l2 = &ListNode{}
		}
		t.Val += l1.Val + l2.Val + carry
		carry = 0
		if t.Val >= 10 {
			t.Val -= 10
			carry = 1
		}

		if l1.Next == nil && l2.Next == nil {
			if carry != 0 {
				t.Next = &ListNode{Val: 1}
			}
			break
		}
		l1 = l1.Next
		l2 = l2.Next
		t.Next = &ListNode{}
		t = t.Next
	}
	return &s
}
