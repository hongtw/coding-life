package main

import (
	"fmt"
)

func isAddOK(x, y int64) bool {
	xSign := (x >> 63) & 1 // x >= 0 ? 0 : 1
	ySign := (y >> 63) & 1 // y >= 0 ? 0 : 1

	neg := xSign*x + xSign*y
	pos := (1-xSign)*x + (1-ySign)*y

	isPosOK := (pos>>63)&1 ^ 1
	isNegOK := ((((xSign + ySign) - (neg>>63)&1) ^ 3) >> 1) & 1

	return (isPosOK & isNegOK) == 1
}

func main() {
	var maxInt64 int64 = 9223372036854775807
	// var minInt64 int64 = -9223372036854775808
	fmt.Println(isAddOK(maxInt64/2, maxInt64/2))
}
