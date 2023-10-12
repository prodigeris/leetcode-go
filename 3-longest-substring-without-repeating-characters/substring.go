package __longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	max := 0
	l := 0
	lookup := make(map[uint8]bool)
	for r := 0; r < len(s); r++ {
		if _, ok := lookup[s[r]]; ok {
			delete(lookup, s[l])
			r--
			l++
			continue
		}

		lookup[s[r]] = true
		if max < r-l+1 {
			max = r - l + 1
		}
	}
	return max
}
