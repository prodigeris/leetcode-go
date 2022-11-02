package leetcode

func romanToInt(s string) int {
	m := make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	p := 0
	sum := 0
	for _, l := range s {
		a := m[string(l)]
		if p < a {
			sum -= 2 * p
			sum += a
		} else {
			sum += a
		}
		p = a
	}

	return sum
}

func RomanToInt(s string) int {
	return romanToInt(s)
}
