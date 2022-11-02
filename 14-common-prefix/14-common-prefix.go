package leetcode

func longestCommonPrefix(strs []string) string {
	c := ""
first:
	for i := 0; i < len(strs[0]); i++ {
		for _, w := range strs {
			if len(w) < i+1 || w[i] != strs[0][i] {
				break first
			}
		}
		c += string(strs[0][i])
	}
	return c
}

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}
