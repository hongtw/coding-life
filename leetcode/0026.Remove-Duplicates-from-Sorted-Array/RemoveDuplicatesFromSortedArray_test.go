package leetcode

import (
	"testing"
)

func Test_RemoveDuplicatesFromSortedArray(t *testing.T) {
	testSuits := []struct {
		name    string
		nums    []int
		answer  int
		valNums []int
	}{
		{"Example 1", []int{1, 1, 2}, 2, []int{1, 2}},
		{"Example 2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{"Example 3", []int{1, 1}, 1, []int{1}},
		{"Example 4", []int{}, 0, []int{}},
	}
	for _, test := range testSuits {
		got := removeDuplicates(test.nums)

		if got != test.answer {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.answer, got)
		}
		for idx, val := range test.valNums {
			if test.nums[idx] != val {
				t.Errorf("[%v][Validate Array] expect '%v', got '%v' ", test.name, test.valNums, test.nums)
				break
			}
		}
	}
}
