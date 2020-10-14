package leetcode

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Ex1", args{[]int{1, 2, 5}, 11}, 3},
		{"Ex2", args{[]int{2}, 3}, -1},
		{"Ex3", args{[]int{1}, 0}, 0},
		{"Ex4", args{[]int{1}, 1}, 1},
		{"Ex5", args{[]int{1}, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
