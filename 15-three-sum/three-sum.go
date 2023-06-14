package _5_three_sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	r := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}

		j, k := i+1, len(nums)-1
		for j < k {
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			if nums[j]+nums[k] == -nums[i] {
				r = append(r, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}

			if nums[j]+nums[k] > -nums[i] {
				k--
			}
			if nums[j]+nums[k] < -nums[i] {
				j++
			}
		}
	}

	return r
}
