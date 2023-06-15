package _1_container_most_water

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		v1, v2 := height[i], height[j]
		s := j - i
		if v1 < v2 {
			s *= v1
			i++
		} else {
			s *= v2
			j--
		}
		if s > max {
			max = s
		}
	}
	return max
}
