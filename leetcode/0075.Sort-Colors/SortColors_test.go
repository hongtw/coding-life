
package leetcode

import (
	"testing"
)

func Test(t *testing.T) {
	testSuits := []struct {
		name string
		nums []int
		ans  []int
	}{
		{"Example 1", []int{2,0,2,1,1,0}, []int{0,0,1,1,2,2}},
	}

	for _, test := range testSuits {
		sortColors(test.nums)
        if !reflect.DeepEqual(test.nums, test.ans) {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, test.nums)
		}

		sortColorsV2(test.nums)
		if !reflect.DeepEqual(test.nums, test.ans) {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.ans, test.nums)
		}
	}
}
