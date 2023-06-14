package _5_three_sum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	r := make([][]int, 0)
	s := make(map[int]bool)
	u := make(map[string]bool)

	for i, v := range nums {
		if _, ok := s[v]; ok {
			continue
		}
		if v > 0 {
			break
		}
		s[v] = true

		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				key := fmt.Sprintf("%d%d%d", v, nums[j], nums[k])
				if v+nums[j]+nums[k] > 0 {
					break
				}
				if _, exists := u[key]; !exists && v+nums[j]+nums[k] == 0 {
					r = append(r, []int{v, nums[j], nums[k]})
					u[key] = true
					break
				}
			}
		}
	}

	return r
}
