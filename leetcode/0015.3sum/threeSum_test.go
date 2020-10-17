package leetcode

import (
	"reflect"
	"testing"
)

func Test_RemoveDuplicatesFromSortedArray(t *testing.T) {
	testSuits := []struct {
		name string
		nums []int
		ans  [][]int
	}{
		{"Example 1", []int{-1, 0, 1, 2, -1, -4}, [][]int{[]int{-1, 0, 1}, []int{-1, -1, 2}}},
	}

	for _, test := range testSuits {
		got := threeSum(test.nums)

		if !reflect.DeepEqual(got, test.ans) {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, got)
		}
	}
}
