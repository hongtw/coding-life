package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testSuits := []struct {
		name string
		nums []int
		ans  []int
	}{
		{"Example 1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"Example 2", []int{2, 1}, []int{2, 1}},
		{"Example 3", []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}, []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0}},
	}

	for _, test := range testSuits {
		moveZeroes(test.nums)

		if !reflect.DeepEqual(test.nums, test.ans) {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, test.nums)
		}
	}
}
