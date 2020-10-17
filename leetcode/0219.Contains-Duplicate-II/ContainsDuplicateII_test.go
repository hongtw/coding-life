package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testSuits := []struct {
		name string
		nums []int
		k    int
		ans  bool
	}{
		{"Example 1", []int{1, 2, 3, 1}, 3, true},
		{"Example 2", []int{1, 0, 1, 1}, 1, true},
		{"Example 3", []int{1, 2, 3, 1, 2, 3}, 2, false},
		{"Example 4", []int{99, 99}, 2, true},
		{"Example 5", []int{1, 2, 1}, 0, false},
	}

	for _, test := range testSuits {
		got := containsNearbyDuplicate(test.nums, test.k)

		if !reflect.DeepEqual(got, test.ans) {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, got)
		}
	}
}
