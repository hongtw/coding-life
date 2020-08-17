package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testSuits := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		ans   []int
	}{
		{"Example 1", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"Example 2", []int{0}, 0, []int{1}, 1, []int{1}},
		{"Example 3", []int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
	}
	for _, test := range testSuits {
		merge(test.nums1, test.m, test.nums2, test.n)

		if !reflect.DeepEqual(test.nums1, test.ans) {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, test.nums1)
		}
	}
}
