package leetcode

import (
	"testing"
)

func Test_removeElement(t *testing.T) {
	testSuits := []struct {
		name    string
		nums    []int
		val     int
		ans     int
		valNums []int
	}{
		{"Example 1", []int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{"Example 2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
	}
	for _, test := range testSuits {
		got := removeElement(test.nums, test.val)

		if got != test.ans {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, got)
		}
		for idx, val := range test.valNums {
			if test.nums[idx] != val {
				t.Errorf("[%v][Validate Array] expect '%v', got '%v' ", test.name, test.valNums, test.nums)
				break
			}
		}
	}
}
