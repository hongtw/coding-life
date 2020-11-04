// https://www.codewars.com/kata/5279f6fe5ab7f447890006a7

package main

import (
	"reflect"
	"testing"
)

func TestPickPeaks(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want PosPeaks
	}{
		// TODO: Add test cases.
		{"Ex1", args{[]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1}}, PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}}},
		{"Ex1", args{[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1}}, PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PickPeaks(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PickPeaks() = %v, want %v", got, tt.want)
			}
		})
	}
}
