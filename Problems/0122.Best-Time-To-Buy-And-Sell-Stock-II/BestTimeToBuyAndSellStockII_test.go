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
		{"Example 1", []int{7, 1, 5, 3, 6, 4}, 7},
		{"Example 2", []int{7, 6, 4, 3, 1}, 0},
		{"Example 3", []int{1, 2, 3, 4, 5}, 4},
	}

	for _, test := range testSuits {
		got := maxProfit(test.nums)

		if got != test.ans {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, got)
		}
	}
}
