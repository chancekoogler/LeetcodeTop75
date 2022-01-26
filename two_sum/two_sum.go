package leetcode

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		desiredValue := target - nums[i]
		if _, ok := m[desiredValue]; ok {
			return []int{i, m[desiredValue]}
		}
		m[nums[i]] = i
	}
	return nil
}
