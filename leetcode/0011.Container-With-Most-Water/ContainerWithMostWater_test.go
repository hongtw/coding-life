package leetcode

import "testing"

func Test_maxArea(t *testing.T) {
	testSuits := []struct {
		name string
		nums []int
		ans  int
	}{
		{"Example 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"Example 2", []int{2, 3, 4, 5, 18, 17, 6}, 17},
	}

	for _, test := range testSuits {
		got := maxArea(test.nums)

		if got != test.ans {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, got)
		}
	}
}
