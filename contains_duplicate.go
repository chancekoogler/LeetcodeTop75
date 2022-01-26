package leetcode

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = i
	}
	return false
}
