package leetcode

import "testing"

func TestContainsDuplicate(t *testing.T) {
	nums := []int{8, 2, 3, 2, 1, 4}
	ans := ContainsDuplicate(nums)
	if ans != true {
		t.Errorf("Got %t. Expected: true", ans)
	}
}
