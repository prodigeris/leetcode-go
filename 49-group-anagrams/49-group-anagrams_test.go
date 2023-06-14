package _9_group_anagrams

import (
	"reflect"
	"testing"
)

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
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
