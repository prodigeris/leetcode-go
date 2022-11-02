package main

//func main() {
//	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 0))
//}
//
//func search(nums []int, target int) int {
//	first := 0
//	last := len(nums) - 1
//	for first <= last {
//		midpoint := (first + last) / 2
//		if nums[midpoint] == target {
//			return midpoint
//		}
//		if nums[midpoint] < target {
//			first = midpoint + 1
//		} else {
//			last = midpoint - 1
//		}
//	}
//
//	return -1
//}

//func search(nums []int, target int) int {
//	return s(nums, target, 0)
//}
//
//func s(nums []int, target int, index int) int {
//	c := len(nums)
//	g := c / 2
//
//	if c == 0 {
//		return -1
//	}
//
//	if nums[g] == target {
//		return index + g
//	}
//
//	if c == 1 {
//		return -1
//	}
//
//	if nums[g] > target {
//		nums = nums[0:g]
//	} else {
//		index += g + 1
//		nums = nums[g+1:]
//	}
//
//	return s(nums, target, index)
//}

//func main() {
//
//	s := []byte("Hannah")
//	reverseString(s)
//	fmt.Println(string(s))
//}

//func reverseString(s []byte) {
//	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
//		fmt.Println(i, j)
//		s[i], s[j] = s[j], s[i]
//	}
//}
