package main

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Ex0", args{"babcde", "acxe"}, 3},
		{"Ex0", args{"babccde", "acxe"}, 3},
		{"Ex0", args{"oxcpqrsvwf", "shmtulqrypy"}, 2},
		{"Ex0", args{"abcba", "abcbcba"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
