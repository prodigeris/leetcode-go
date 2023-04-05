package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			break
		}

		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
		}
	}

	return digits
}
