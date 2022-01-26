package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var testNums = []struct {
		a, b, c, d int
		want       int
	}{
		{2, 11, 7, 15, 22},
		{3, 6, 12, 11, 18},
		{1, 3, 2, 5, 5},
		{1, 10, 25, 50000, 35},
	}
	for i, val := range testNums {
		testName := fmt.Sprintf("Test %d. Using %d, %d, %d, %d", i, val.a, val.b, val.c, val.d)
		t.Run(testName, func(t *testing.T) {
			nums := []int{val.a, val.b, val.c, val.d}
			ans := TwoSum(nums, val.want)
			if nums[ans[0]]+nums[ans[1]] != val.want {
				t.Errorf("got %d, want %d", ans, val.want)
			}

		})
	}
}
