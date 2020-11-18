// https://www.codewars.com/kata/5a24a35a837545ab04001614
package main

import (
	"fmt"
	"math"
	"strings"
)

type cipherFunc func(s string) string

func buildCipherTable(n int, doEncode bool) (table []int) {
	/*
		isEncode
		3: []int{ 0, 4, 1, 7, 8, 5, 3, 6, 2}
		4: []int{ 0, 4, 8, 1, 11, 12, 13, 5, 7, 15, 14, 9, 3, 10, 6, 2}
		5: []int{ 0, 4, 8, 12, 1, 15, 16, 20, 17, 5, 11, 23, 24, 21, 9, 7, 19, 22, 18, 13, 3, 14, 10, 6, 2}

	*/
	arr := make([]int, n*n)
	base := 0
	round := 0
	for i := 0; i < n*n; i++ {
		sin := int(math.Round(math.Sin(math.Pi / 2 * float64(i))))
		cos := int(math.Round(math.Cos(math.Pi / 2 * float64(i))))
		row := (sin+cos)*round + sin*(i-base)/4 + n*(1-sin-cos)/2 + (sin+cos-1)/2
		col := (cos-sin)*round + cos*(i-base)/4 + n*(sin-cos+1)/2 + (cos-sin-1)/2
		position := row*n + col
		// fmt.Println("*****", i, sin, cos, row, col)
		if i%4 == 3 && row == round+1 {
			round++
			base = i + 1
		}
		if doEncode {
			arr[position] = i
		} else {
			arr[i] = position
		}

	}
	fmt.Printf("View Table(%v) %v\n", n, arr)
	return arr
}

func cipherHelper(s string, doEncode bool) string {
	total := len(s)
	sideLen := int(math.Ceil(math.Sqrt(float64(len(s)))))
	table := buildCipherTable(sideLen, doEncode)
	ret := ""
	for i := 0; i < sideLen*sideLen; i++ {
		if stringIdx := table[i]; stringIdx < total {
			ret += string(s[stringIdx])
		} else {
			ret += " "
		}
	}
	if !doEncode {
		ret = strings.TrimSpace(ret)
	}
	return ret
}

var encode cipherFunc = func(s string) string {
	return cipherHelper(s, true)
}

var decode cipherFunc = func(s string) string {
	return cipherHelper(s, false)
}

var InterlacedSpiralCipher = map[string]cipherFunc{"encode": encode, "decode": decode}

func main() {
	fmt.Println(encode("012345678"))
	fmt.Println(decode("041785362"))
	fmt.Println(decode("Rntodomiimuea  m"))

	buildCipherTable(3, true)
	buildCipherTable(3, false)
	buildCipherTable(4, true)
	buildCipherTable(5, true)

}
