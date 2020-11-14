// https://www.codewars.com/kata/5a24a35a837545ab04001614
package main

import (
	"fmt"
	"math"
	"strings"
)

type cipherFunc func(s string) string

func buildIndexTable(n int) (table []int) {
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

var encode cipherFunc = func(s string) string {
	total := len(s)
	borderLen := int(math.Ceil(math.Sqrt(float64(len(s)))))
	table := buildIndexTable(borderLen)
	ret := ""
	for i := 0; i < borderLen*borderLen; i++ {
		stringIdx := table[i]
		if stringIdx < total {
			ret += string(s[stringIdx])
		} else {
			ret += " "
		}
	}
	return ret
}

var decode cipherFunc = func(s string) string {
	total := len(s)
	borderLen := int(math.Ceil(math.Sqrt(float64(len(s)))))
	table := buildIndexTable(borderLen)
	ret := make([]byte, total)
	for i := 0; i < borderLen*borderLen; i++ {
		if i < total {
			ret[table[i]] = s[i]
		}
	}
	return strings.TrimSpace(string(ret))
}

var InterlacedSpiralCipher = map[string]cipherFunc{"encode": encode, "decode": decode}

func main() {
	fmt.Println(encode("012345678"))
	fmt.Println(decode("041785362"))
	buildIndexTable(3)
	buildIndexTable(4)
	buildIndexTable(5)
}
