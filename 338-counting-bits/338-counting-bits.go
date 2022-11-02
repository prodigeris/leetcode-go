package leetcode

func countBits(n int) []int {
	r := make([]int, n+1)
	r[0] = 0
	for i := 1; i <= n; i++ {
		r[i] = i%2 + r[i/2]
	}
	return r
}

// 0  - 0
// 1  -
// 1 % 2 = 1

// 2  -
// 2 % 2 = 0,
// 1 % 2 = 1

// 3  -
// 3 % 2 = 1
// 1 % 2 = 1

// 4  -
// 4 % 2 = 0
// 2 % 2 = 0
// 1 % 2 = 1

// 5  -
// 5 % 2 = 1
// 2 % 2 = 0
// 1 % 2 = 1

// result: [0, 1, 2, 1, 2]
