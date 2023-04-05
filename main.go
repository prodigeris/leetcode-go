package main

import (
	"fmt"
	leetcode "leetcode/2-add-two-numbers"
)

func main() {
	l1 := intToListNode(9999999)
	l2 := intToListNode(9999)
	fmt.Println(listNodeToString(leetcode.AddTwoNumbers(l1, l2)))
}

func intToListNode(i int) *leetcode.ListNode {
	var l *leetcode.ListNode
	i = reverseNumber(i)
	for i > 0 {
		l = &leetcode.ListNode{
			Val:  i % 10,
			Next: l,
		}
		i /= 10
	}
	return l
}

func listNodeToString(n *leetcode.ListNode) string {
	s := ""
	for {
		s += fmt.Sprintf("%d", n.Val)
		if n.Next == nil {
			break
		}
		n = n.Next
	}
	return s
}

func reverseNumber(num int) int {

	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}
