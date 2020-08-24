
package leetcode

import (
	"testing"
)

func Test(t *testing.T) {
	testSuits := []struct {
		name string
		nums []int
		k    int
		ans  int
	}{
		{"Example 1", []int{6, 3, 5, 7, 2, 3, 3, 8, 2, 4}, 2, 5},
	}

	for _, test := range testSuits {
		got := findPairs(test.nums, test.k)

		if got != test.ans {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, got)
		}
	}
}
