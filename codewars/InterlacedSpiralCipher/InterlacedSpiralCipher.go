// https://www.codewars.com/kata/5a24a35a837545ab04001614
package main

import (
	"fmt"
	"math"
	"strings"
)

type cipherFunc func(s string) string

// first implement
func buildIndexTableLegacy(n int) (table []int) {
	/*
		3: []int{ 0, 4, 1, 7, 8, 5, 3, 6, 2}
		4: []int{ 0, 4, 8, 1, 11, 12, 13, 5, 7, 15, 14, 9, 3, 10, 6, 2}
		5: []int{ 0, 4, 8, 12, 1, 15, 16, 20, 17, 5, 11, 23, 24, 21, 9, 7, 19, 22, 18, 13, 3, 14, 10, 6, 2}
	*/
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	base := 0
	round := 0
	for i := 0; i < n*n; i++ {
		order := i % 4
		if order == 0 {
			arr[0+round][(i-base)/4+round] = i
		} else if order == 1 {
			arr[(i-base)/4+round][n-1-round] = i
		} else if order == 2 {
			arr[n-1-round][n-(i-base)/4-1-round] = i
		} else {
			arr[n-(i-base)/4-1-round][0+round] = i
			if n-(i-base)/4-1-round == round+1 {
				round++
				base = i + 1
			}
		}
	}
	fmt.Printf("View Table(%v) %v\n", n, arr)
	for _, row := range arr {
		table = append(table, row...)
	}
	return table
}

func buildIndexTable(n int, doEncode bool) (table []int) {
	/*
		isEncode
		3: []int{ 0, 4, 1, 7, 8, 5, 3, 6, 2}
		4: []int{ 0, 4, 8, 1, 11, 12, 13, 5, 7, 15, 14, 9, 3, 10, 6, 2}
		5: []int{ 0, 4, 8, 12, 1, 15, 16, 20, 17, 5, 11, 23, 24, 21, 9, 7, 19, 22, 18, 13, 3, 14, 10, 6, 2}

	*/
	arr := make([]int, n*n)
	base := 0
	round := 0
	var r, c int
	for i := 0; i < n*n; i++ {
		order := i % 4
		if order == 0 {
			r = 0 + round
			c = (i-base)/4 + round
		} else if order == 1 {
			r = (i-base)/4 + round
			c = n - 1 - round
		} else if order == 2 {
			r = n - 1 - round
			c = n - (i-base)/4 - 1 - round
		} else {
			r = n - (i-base)/4 - 1 - round
			c = 0 + round
			if r == round+1 {
				round++
				base = i + 1
			}
		}
		if doEncode {
			arr[r*n+c] = i
		} else {
			arr[i] = r*n + c
		}

	}
	fmt.Printf("View Table(%v) %v\n", n, arr)
	return arr
}

func cipherHelper(s string, doEncode bool) string {
	total := len(s)
	sideLen := int(math.Ceil(math.Sqrt(float64(len(s)))))
	table := buildIndexTable(sideLen, doEncode)
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

	buildIndexTable(3, true)
	buildIndexTable(4, true)
	buildIndexTable(5, true)

}
