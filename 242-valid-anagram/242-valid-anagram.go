package leetcode

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	av := make(map[int32]int)
	for _, c := range s {
		av[c] += 1
	}

	fmt.Println(av)
	for _, c := range t {
		if v, ok := av[c]; ok && v > 0 {
			av[c] -= 1
		} else {
			return false
		}
	}

	return true
}
