package _1_container_most_water

import (
	"reflect"
	"testing"
)

func Test_maxArea(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"Example 2", args{[]int{1, 1}}, 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := maxArea(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
