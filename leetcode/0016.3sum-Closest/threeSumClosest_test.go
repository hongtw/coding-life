package leetcode

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{-1, 2, 1, -4}, 1}, 2},
		{"Example 2", args{[]int{1, 1, 1, 0}, -100}, 2},
		{"Example 3", args{[]int{0, 2, 1, -3}, 1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
