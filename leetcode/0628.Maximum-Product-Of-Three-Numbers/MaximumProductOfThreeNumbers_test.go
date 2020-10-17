package leetcode

import "testing"

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example-1", args{[]int{1, 2, 3}}, 6},
		{"Example-2", args{[]int{1, 2, 3, 4}}, 24},
		{"Example-3", args{[]int{-3, -5, 2, 3}}, 45},
		{"Example-4", args{[]int{-1, -2, -3}}, -6},
		{"Example-5", args{[]int{-1, -2, -3, -4}}, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
