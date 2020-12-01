package main

import "testing"

func Test_rgb(t *testing.T) {
	type args struct {
		r int
		g int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"ex", args{0, 0, -20}, "000000"},
		{"ex", args{300, 255, 255}, "FFFFFF"},
		{"ex", args{173, 255, 47}, "ADFF2F"},
		{"ex", args{258, 206, -4}, "FFCE00"},
		{"ex", args{9, 165, 222}, "09A5DE"},
		{"ex", args{256, 69, 277}, "FF45FF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rgb(tt.args.r, tt.args.g, tt.args.b); got != tt.want {
				t.Errorf("rgb() = %v, want %v", got, tt.want)
			}
		})
	}
}
