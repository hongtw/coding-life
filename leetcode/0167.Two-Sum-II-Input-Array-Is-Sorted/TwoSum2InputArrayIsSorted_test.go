package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testSuits := []struct {
		name   string
		nums   []int
		target int
		ans    []int
	}{
		{"Example 1", []int{2, 7, 11, 15}, 9, []int{1, 2}},
	}
	for _, test := range testSuits {
		got := twoSum(test.nums, test.target)

		if !reflect.DeepEqual(got, test.ans) {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, got)
		}
	}
}
