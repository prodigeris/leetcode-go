package leetcode

import (
	"fmt"
	"testing"
)

func Test_firstBadVersion(t *testing.T) {
	t.Run("wow", func(t *testing.T) {
		t.Parallel()
		fmt.Println("rezultatas", firstBadVersion(5))
	})
}
