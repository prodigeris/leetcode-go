package _9_group_anagrams

func groupAnagrams(strs []string) [][]string {
	r := make([][]string, 0)
	k := 0
	used := make(map[string]bool)
	for i, str := range strs {
		if used[str] {
			continue
		}
		r = append(r, []string{str})
		used[str] = true
		for j, str2 := range strs {
			if i == j {
				continue
			}
			if isAnagram(str, str2) {
				r[k] = append(r[k], str2)
				used[str2] = true
			}
		}
		k++
	}
	return r
}

func isAnagram(str string, str2 string) bool {
	if len(str) == 0 && len(str2) == 0 {
		return true
	}
	if len(str) != len(str2) {
		return false
	}
	for i, v := range str2 {
		if str[0] == uint8(v) {
			return isAnagram(str[1:], str2[:i]+str2[i+1:])
		}
	}
	return false
}
