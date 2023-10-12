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
