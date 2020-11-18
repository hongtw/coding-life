package main

import (
	"fmt"
	"math"
)

// first implement
func buildIndexTableV1(n int) (table []int) {
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

func buildIndexTableV2(n int, doEncode bool) (table []int) {
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

func buildCipherTableFinal(n int, doEncode bool) (table []int) {
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

// func main() {
// 	buildIndexTableV1(3)
// 	buildIndexTableV1(4)
// 	buildIndexTableV1(5)

// 	buildIndexTableV2(3, true)
// 	buildIndexTableV2(4, true)
// 	buildIndexTableV2(5, true)
// }
