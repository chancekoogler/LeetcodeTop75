package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	ans := MaxProfit(nums)
	if ans != 5 {
		t.Errorf("Expected 5, got %v", ans)
	}
}
