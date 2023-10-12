package __longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	max := 1
	lookup := make([]map[uint8]bool, 0)
	for l := 0; l < len(s)-1; l++ {
		r := l + 1
		lookup = append(lookup, map[uint8]bool{s[l]: true})
		for _, look := range lookup {
			if _, ok := look[s[r]]; !ok {
				look[s[r]] = true
			} else {
				if max < len(look) {
					max = len(look)
				}
				lookup = lookup[1:]
			}
		}
	}
	if len(lookup) > 0 && max < len(lookup[0]) {
		return len(lookup[0])
	}
	return max
}
