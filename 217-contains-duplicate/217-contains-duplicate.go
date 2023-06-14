package _17_contains_duplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, v := range nums {
		if _, exists := m[v]; exists {
			return true
		} else {
			m[v] = true
		}
	}
	return false
}
