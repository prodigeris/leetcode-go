package _9_group_anagrams

import (
	"reflect"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	type args struct {
		s  string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{"alus", "sula"}, true},
		{"Example 2", args{"arnas", "varnas"}, false},
		{"Example 3", args{"bat", "tab"}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isAnagram(tt.args.s, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isAnagram(%v, %v) = %v, want %v", tt.args.s, tt.args.s2, got, tt.want)
			}
		})
	}
}

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"Example 1", args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{"Example 2", args{[]string{""}}, [][]string{{""}}},
		{"Example 3", args{[]string{"a"}}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := groupAnagrams(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
