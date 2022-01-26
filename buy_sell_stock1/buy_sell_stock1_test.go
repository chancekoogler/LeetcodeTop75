package leetcode

import "testing"

type testNums []struct {
	nums []int
	want int
}

func TestMaxProfit(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	ans := MaxProfit(nums)
	if ans != 5 {
		t.Errorf("Expected 5, got %v", ans)
	}
}

func TestMaxProfits(t *testing.T) {
	var testNums = []struct {
		nums []int
		want int
	}{
		{
			nums: []int{7, 1, 5, 3, 6, 4},
			want: 5,
		},
		{
			nums: []int{7, 2, 5, 3, 6, 4},
			want: 4,
		},
		{
			nums: []int{7, 3, 5, 3, 6, 4},
			want: 3,
		},
	}

	// for i := range testNums {

	// }
}
