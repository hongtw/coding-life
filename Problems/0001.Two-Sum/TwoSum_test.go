package leetcode

import (
	"reflect"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	testSuits := []struct {
		name   string
		nums   []int
		target int
		answer []int
	}{
		{"Example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Example 2", []int{4, 6, 2, 1}, 8, []int{1, 2}},
	}
	for _, test := range testSuits {
		got := twoSum(test.nums, test.target)

		if !reflect.DeepEqual(got, test.answer) {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.answer, got)
		}
	}
}
