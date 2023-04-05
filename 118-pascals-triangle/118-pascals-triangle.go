package leetcode

func Generate(numRows int) [][]int {
	return generate(numRows)
}

func generate(numRows int) [][]int {
	r := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		r[i] = make([]int, i+1)
		r[i][0] = 1
		r[i][i] = 1

		for j := 1; j < i; j++ {
			r[i][j] = r[i-1][j-1] + r[i-1][j]
		}
	}

	return r
}
