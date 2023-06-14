package _5_three_sum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{name: "Example 2", args: args{[]int{0, 1, 1}}, want: [][]int{}},
		{name: "Example 3", args: args{[]int{0, 0, 0}}, want: [][]int{{0, 0, 0}}},
		{name: "Example 4", args: args{[]int{0, 0, 0, 0}}, want: [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := threeSum(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
