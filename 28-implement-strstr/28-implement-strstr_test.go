package leetcode

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"hello", "ll"}, 2},
		{"Example 2", args{"aaaaa", "bba"}, -1},
		{"Example 3", args{"aaaaa", ""}, 0},
		{"Example 4", args{"aaa", "aaaa"}, -1},
		{"Example 5", args{"mississippi", "issip"}, 4},
		{"Example 6", args{"mississippi", "issipi"}, -1},
		{"Example 7", args{"mississippi", "pi"}, 9},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
