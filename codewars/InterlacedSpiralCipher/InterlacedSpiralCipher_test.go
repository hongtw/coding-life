// https://www.codewars.com/kata/5a24a35a837545ab04001614

package main

import "testing"

func Test_cipherHelper(t *testing.T) {
	type args struct {
		s        string
		doEncode bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Ex-encode-1", args{"Romani ite domum", true}, "Rntodomiimuea  m"},
		{"Ex-decode-1", args{"Rntodomiimuea  m", false}, "Romani ite domum"},
		{"Ex-encode-2", args{"Sic transit gloria mundi", true}, "Stsgiriuar i ninmd l otac"},
		{"Ex-decode-3", args{"Stsgiriuar i ninmd l otac", false}, "Sic transit gloria mundi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cipherHelper(tt.args.s, tt.args.doEncode); got != tt.want {
				t.Errorf("cipherHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}
