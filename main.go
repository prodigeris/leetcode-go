package main

import (
	"fmt"
	leetcode "leetcode/20-valid-parentheses"
)

func main() {
	fmt.Println(leetcode.IsValid("]"))
}

//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func main() {
//	//c := ListNode{Val: 4, Next: nil}
//	//b := ListNode{Val: 2, Next: &c}
//	//list1 := ListNode{Val: 1, Next: &b}
//
//	c2 := ListNode{Val: 4, Next: nil}
//	b2 := ListNode{Val: 3, Next: &c2}
//	list2 := ListNode{Val: 1, Next: &b2}
//
//	r := mergeTwoLists(nil, &list2)
//
//	for true {
//		fmt.Println(r.Val)
//		r = r.Next
//		if r == nil {
//			break
//		}
//	}
//}
//
///**
// * Definition for singly-linked list.
// * type ListNode struct {
// *     Val int
// *     Next *ListNode
// * }
// */
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	var left, right int
//	var result, last *ListNode
//	for true {
//		if list1 == nil {
//			if last == nil {
//				last = list2
//				result = last
//			} else {
//				last.Next = list2
//			}
//			break
//		}
//
//		if list2 == nil {
//			if last == nil {
//				last = list1
//				result = last
//			} else {
//				last.Next = list1
//			}
//			break
//		}
//		left = list1.Val
//		right = list2.Val
//		if left < right {
//			if last != nil {
//				last.Next = list1
//			}
//			last = list1
//			list1 = list1.Next
//		} else {
//			if last != nil {
//				last.Next = list2
//			}
//			last = list2
//			list2 = list2.Next
//		}
//		if result == nil {
//			result = last
//		}
//	}
//	return result
//}
