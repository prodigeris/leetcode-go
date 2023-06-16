package search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	first, last := 0, len(matrix)-1
	for first <= last {
		midpoint := (first + last) / 2
		if matrix[midpoint][0] == target {
			return true
		}
		if matrix[midpoint][0] < target {
			first = midpoint + 1
		} else {
			last = midpoint - 1
		}
	}
	r := last
	if r < 0 {
		r = 0
	}

	first, last = 0, len(matrix[r])-1
	for first <= last {
		midpoint := (first + last) / 2
		if matrix[r][midpoint] == target {
			return true
		}
		if matrix[r][midpoint] < target {
			first = midpoint + 1
		} else {
			last = midpoint - 1
		}
	}

	return false
}

func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}
