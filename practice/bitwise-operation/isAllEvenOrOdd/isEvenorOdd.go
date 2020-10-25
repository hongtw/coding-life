package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isAllEven(x int64) bool {
	/*
		"01010101....01" -> true
	*/
	return (x + x<<1) == -1
}

func isAllOdd(x int64) bool {
	/*
		"1010101....010" -> true
	*/
	return (x + x<<1) == -2
}

func main() {
	allEvenStr := strings.Repeat("01", 32)
	allEvenNum, _ := strconv.ParseInt(allEvenStr, 2, 64)
	allOddNum := allEvenNum << 1
	allOddStr := fmt.Sprintf("%b", uint(allOddNum))

	fmt.Println(allEvenStr, "is All Even: ", isAllEven(allEvenNum))
	fmt.Println(allOddStr, "is All Odd:  ", isAllOdd(allOddNum))
}
