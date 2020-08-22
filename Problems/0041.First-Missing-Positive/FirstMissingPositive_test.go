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
		{"Example 1", []int{1, 2, 0}, 3},
		{"Example 2", []int{3, 4, -1, 1}, 2},
		{"Example 3", []int{7, 8, 9, 11, 12}, 1},
	}

	for _, test := range testSuits {
		got := firstMissingPositive(test.nums)

		if got != test.ans {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, got)
		}
	}
}
