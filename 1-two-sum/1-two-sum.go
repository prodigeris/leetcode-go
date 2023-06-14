package __two_sum

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j, v2 := range nums {
			if i == j {
				continue
			}
			if v+v2 == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
