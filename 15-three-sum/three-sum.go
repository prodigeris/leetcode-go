package _5_three_sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	r := make([][]int, 0)
	s := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		v1 := nums[i]
		if _, ok := s[v1]; ok {
			continue
		}
		if v1 > 0 {
			break
		}
		s[v1] = true

		j, k := i+1, len(nums)-1
		for j < k {
			v2, v3 := nums[j], nums[k]

			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			if v2+v3 > -v1 {
				k--
			}
			if v2+v3 < -v1 {
				j++
			}

			if v2+v3 == -v1 {
				r = append(r, []int{v1, v2, v3})
				j++
				k--
			}
		}
	}

	return r
}
