package main

import "testing"

func Test_searchSubstr(t *testing.T) {
	type args struct {
		fullText     string
		searchText   string
		allowOverlap bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Ex1", args{"aa_bb_cc_dd_bb_e", "bb", true}, 2},
		{"Ex2", args{"aaabbbcccc", "bbb", true}, 1},
		{"Ex3", args{"aaa", "aa", true}, 2},
		{"Ex4", args{"aaa", "aa", false}, 1},
		{"Ex5", args{"", "aa", true}, 0},
		{"Ex6", args{"aaa", "", false}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchSubstr(tt.args.fullText, tt.args.searchText, tt.args.allowOverlap); got != tt.want {
				t.Errorf("searchSubstr() = %v, want %v", got, tt.want)
			}
		})
	}
}
