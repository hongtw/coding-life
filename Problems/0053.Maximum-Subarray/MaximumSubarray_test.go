package leetcode

import (
	"testing"
)

func Test(t *testing.T) {
	testSuits := []struct {
		name string
		nums []int
		ans  int
	}{
		{"Example 1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"Example 2", []int{-1}, -1},
	}

	for _, test := range testSuits {
		got := maxSubArray(test.nums)

		if got != test.ans {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, got)
		}
	}
}
