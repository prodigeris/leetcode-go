package main

import "fmt"

func subsets(nums []int) [][]int {
	var r [][]int

	dfs(nums, &r, []int{}, 0)

	return r
}

func dfs(nums []int, r *[][]int, c []int, index int) {
	if index == len(nums) {
		copied := make([]int, len(c))
		copy(copied, c)
		*r = append(*r, copied)
		return
	}

	dfs(nums, r, c, index+1)
	dfs(nums, r, append(c, nums[index]), index+1)
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
