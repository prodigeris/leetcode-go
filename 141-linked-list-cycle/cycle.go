package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Hashmap solution
func hasCycle(head *ListNode) bool {
	nodes := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := nodes[head]; ok {
			return true
		}
		nodes[head] = true
		head = head.Next
	}
	return false
}

func hasCycleTortoise(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
