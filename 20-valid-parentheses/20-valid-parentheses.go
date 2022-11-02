package leetcode

func isValid(s string) bool {
	opened := make([]int, 0)
	o1, o2, o3 := 0, 0, 0
	for _, s := range s {
		switch string(s) {
		case "(":
			o1++
			opened = append(opened, 1)
		case ")":
			o1--
			if len(opened) <= 0 || opened[len(opened)-1] != 1 {
				return false
			}
			opened = opened[:len(opened)-1]
		case "{":
			o2++
			opened = append(opened, 2)
		case "}":
			o2--
			if len(opened) <= 0 || opened[len(opened)-1] != 2 {
				return false
			}
			opened = opened[:len(opened)-1]
		case "[":
			o3++
			opened = append(opened, 3)
		case "]":
			o3--
			if len(opened) <= 0 || opened[len(opened)-1] != 3 {
				return false
			}
			opened = opened[:len(opened)-1]
		}
	}
	return o1 == 0 && o2 == 0 && o3 == 0
}

func IsValid(s string) bool {
	return isValid(s)
}
