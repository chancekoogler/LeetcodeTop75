package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 11, 7, 15}
	target := 9
	ans := TwoSum(nums, target)
	if ans[0] != 0 && ans[1] != 2 {
		t.Errorf("TwoSum([2,11,7,15], 9): got: %d; want {0,2}", ans)
	}
}
