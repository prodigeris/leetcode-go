package main

func combinationSum(candidates []int, target int) [][]int {
	var r [][]int

	dfs(&r, candidates, []int{}, 0, target)

	return r
}

func dfs(r *[][]int, candidates []int, curr []int, index int, target int) {
	s := sum(curr)
	if s == target {
		copied := make([]int, len(curr))
		copy(copied, curr)
		*r = append(*r, copied)
		return
	}
	if len(candidates) == index || s > target {
		return
	}

	dfs(r, candidates, curr, index+1, target)
	dfs(r, candidates, append(curr, candidates[index]), index, target)
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}
