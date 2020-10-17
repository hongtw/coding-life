package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testSuits := []struct {
		name string
		nums []int
		ans  bool
	}{
		{"Example 1", []int{1, 2, 3, 1}, true},
		{"Example 2", []int{1, 2, 3, 4}, false},
		{"Example 3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, test := range testSuits {
		got := containsDuplicate(test.nums)

		if !reflect.DeepEqual(got, test.ans) {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, got)
		}
	}
}
