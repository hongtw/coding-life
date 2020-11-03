package main

import (
	"fmt"
)

func isAllEvenOnes(x int32) bool {
	/*
		0111010101110101
				01110101
		----------------
				01110101
					0111
		----------------
					0101
					  01
		-----------------
		              01
	*/
	x16 := (x >> 16) & x
	x8 := (x16 >> 8) & x16
	x4 := (x8 >> 4) & x8
	x2 := (x4 >> 2) & x4
	return x2&1 == 1
}

func isAllOddOnes(x int32) bool {
	return isAllEvenOnes(x >> 1)
}

func main() {
	allEvenNum0 := int32(0b01010101010101010101010101010101)
	allEvenNum1 := int32(0b01110101010101011101010101011101)
	allOddNum0 := ^int32(0b01010101010101010101010101010101)

	fmt.Println(fmt.Sprintf("%032b", uint(allEvenNum0)), "is Ones All Even: ", isAllEvenOnes(allEvenNum0))
	fmt.Println(fmt.Sprintf("%032b", uint(allEvenNum1)), "is Ones All Even: ", isAllEvenOnes(allEvenNum1))
	fmt.Println(fmt.Sprintf("%032b", uint(allOddNum0)), "is Ones All Odd: ", isAllOddOnes(allOddNum0))
}
