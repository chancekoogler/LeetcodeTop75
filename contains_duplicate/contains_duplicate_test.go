package leetcode

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	nums := []int{8, 2, 3, 2, 1, 4}
	ans := ContainsDuplicate(nums)
	if !ans {
		t.Errorf("Got %t. Expected: true", ans)
	}
}

func TestContainsDuplication(t *testing.T) {
	testNums := [][]int{
		{1, 2, 3, 4, 1},
		{3, 4, 5, 6, 7, 3},
		{2, 3, 5, 7, 8, 8},
		{600, 601, 600},
	}
	// fmt.Println(testNums)
	for i := range testNums {
		testName := fmt.Sprintf("Test %d", i)
		t.Run(testName, func(t *testing.T) {
			ans := ContainsDuplicate(testNums[i])
			if !ans {
				t.Errorf("Got %t. Expected: true", ans)
			}
		})
	}
}
