package _5_three_sum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	r := make([][]int, 0)
	u := make(map[string]bool)
	for i := 0; i < len(nums)-2; i++ {
		for j := 1; j < len(nums)-1; j++ {
			for z := 2; z < len(nums); z++ {
				if i == j || j == z || i == z {
					continue
				}
				v1, v2, v3 := nums[i], nums[j], nums[z]
				keys := []int{v1, v2, v3}
				sort.Ints(keys)
				key := fmt.Sprintf("%d%d%d", keys[0], keys[1], keys[2])
				if v1+v2+v3 == 0 {
					if _, exist := u[key]; !exist {
						r = append(r, []int{v1, v2, v3})
						u[key] = true
					}
				}
			}
		}
	}
	return r
}
