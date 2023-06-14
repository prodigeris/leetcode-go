package __two_sum

import (
	"reflect"
	"testing"
)

func Test_twoSums(t *testing.T) {
	type args struct {
		n []int
		t int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"Example 2", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"Example 3", args{[]int{3, 3}, 6}, []int{0, 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := twoSum(tt.args.n, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
