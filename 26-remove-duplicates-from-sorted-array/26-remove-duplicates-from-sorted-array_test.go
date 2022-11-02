package leetcode

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)

	t.Run("wow", func(t *testing.T) {
		t.Parallel()
		if k != 5 {
			t.Errorf("k, gavom %v, norim %v", k, 5)
		}
		for i := 0; i < 5; i++ {
			if i != nums[i] {
				t.Errorf("i, gavom %v, norim %v", nums[i], i)
			}
		}
	})

	fmt.Println(nums)
}
