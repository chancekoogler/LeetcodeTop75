package sum

import "testing"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, _ := range nums {
		desiredValue := target - nums[i]
		if _, ok := m[desiredValue]; ok {
			return []int{i, m[desiredValue]}
		}
		m[nums[i]] = i
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 11, 7, 15}
	target := 9
	ans := twoSum(nums, target)
	if ans[0] != 0 && ans[1] != 2 {
		t.Errorf("TwoSum([2,11,7,15], 9): got: %d; want {0,2}", ans)
	}
}
