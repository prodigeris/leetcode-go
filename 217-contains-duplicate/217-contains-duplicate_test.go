package _17_contains_duplicate

import (
	"reflect"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{[]int{1, 2, 3, 1}}, true},
		{"Example 2", args{[]int{1, 2, 3, 4}}, false},
		{"Example 3", args{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := containsDuplicate(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
