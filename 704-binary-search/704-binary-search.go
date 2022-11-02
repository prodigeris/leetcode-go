package leetcode

func search(nums []int, target int) int {
	first := 0
	last := len(nums) - 1
	for first <= last {
		midpoint := (first + last) / 2
		if nums[midpoint] == target {
			return midpoint
		}
		if nums[midpoint] < target {
			first = midpoint + 1
		} else {
			last = midpoint - 1
		}
	}

	return -1
}
