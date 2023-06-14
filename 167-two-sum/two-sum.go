package _67_two_sum

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1

	for nums[i]+nums[j] != target {
		if nums[i]+nums[j] > target {
			j--
		}
		if nums[i]+nums[j] < target {
			i++
		}
	}
	return []int{i + 1, j + 1}
}
