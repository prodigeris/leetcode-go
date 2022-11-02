package leetcode

func removeDuplicates(nums []int) int {
	c := 0
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			c++
		} else {
			nums[i-c] = nums[i]
		}
	}
	return n - c
}

//func removeDuplicates(nums []int) int {
//	k := len(nums)
//	p := -101
//
//	for i, num := range nums {
//		if p == num {
//			nums[i] = -101
//			k--
//		}
//		p = num
//	}
//
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == -101 {
//			f := false
//			for a := i + 1; a < len(nums); a++ {
//				if nums[a] != -101 {
//					nums[i], nums[a] = nums[a], nums[i]
//					f = true
//					break
//				}
//			}
//			if !f {
//				break
//			}
//		}
//	}
//	return k
//}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
}
