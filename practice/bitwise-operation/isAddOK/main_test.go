package main

import "testing"

var maxInt64 int64 = 9223372036854775807
var minInt64 int64 = -9223372036854775808

func Test_isAddOK(t *testing.T) {
	type args struct {
		x int64
		y int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Ex0", args{13, -22}, true},
		{"Ex0", args{0, 0}, true},
		{"Ex0", args{1, 1}, true},
		{"Ex0", args{-1, -1}, true},
		{"Ex0", args{minInt64, maxInt64}, true},
		{"Ex0", args{maxInt64, 0}, true},
		{"Ex0", args{minInt64, 0}, true},

		{"Ex1", args{maxInt64, 1}, false},
		{"Ex2", args{maxInt64, -1}, true},
		{"Ex3", args{maxInt64, maxInt64}, false},
		{"Ex4", args{maxInt64 / 2, maxInt64 / 2}, true},
		{"Ex5", args{maxInt64/2 + 1, maxInt64/2 + 1}, false},

		{"Ex6", args{minInt64, 1}, true},
		{"Ex7", args{minInt64, -1}, false},
		{"Ex8", args{minInt64, minInt64}, false},
		{"Ex9", args{minInt64 / 2, minInt64 / 2}, true},
		{"Ex10", args{minInt64 / 2, minInt64/2 - 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAddOK(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isAddOK() = %v, want %v", got, tt.want)
			}
		})
	}
}
