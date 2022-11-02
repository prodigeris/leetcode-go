package leetcode

func strStr(haystack string, needle string) int {
	c := 0
	s := -1
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		v := haystack[i]
		if v == needle[c] {
			if c == 0 {
				s = i
			}
			c++
		} else if c != 0 {
			i = s
			c = 0
		}
		if len(needle) <= c {
			break
		}
	}
	if len(needle) == c {
		return s
	}
	return -1
}
