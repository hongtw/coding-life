package main

import (
	"fmt"
)

// Comb : Combinations, reference to https://docs.python.org/3.1/library/itertools.html?highlight=combinations#itertools.combinations
func Comb(ls []int, r int) [][]int {
	length := len(ls)
	indices := make([]int, r)
	reversedIndices := make([]int, r)
	for i := range indices {
		indices[i] = i
		reversedIndices[r-i-1] = i
	}
	result := [][]int{ls[:r]}
	for {
		var i int
		forLoopAll := true
		for _, i = range reversedIndices {
			if indices[i] != i+length-r {
				forLoopAll = false
				break
			}
		}
		if forLoopAll {
			return result
		}
		indices[i]++
		for j := i + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1

		}
		subRes := make([]int, r)
		for k, idx := range indices {
			subRes[k] = ls[idx]
		}
		result = append(result, subRes)
	}
}

// Sum slice of int
func Sum(nums []int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return
}

// ChooseBestSumBruteForce : Brute force
func ChooseBestSumBruteForce(t, k int, ls []int) int {
	if len(ls) < k {
		return -1
	}
	combinations := Comb(ls, k)
	dist := t
	res := -1

	for _, comb := range combinations {
		tmp := Sum(comb)
		if tmp == t {
			return t
		} else if tmp < t && t-tmp <= dist {
			dist = t - tmp
			res = tmp
		}
	}
	return res
}

// ChooseBestSumRecur : Recursive
func ChooseBestSumRecur(t, k int, ls []int) int {
	best := -1
	if len(ls) < k || k < 1 {
		return best
	}
	for i, n := range ls {
		if k > 1 {
			tmpBest := ChooseBestSumRecur(t-n, k-1, ls[i+1:])
			if tmpBest < 0 {
				continue
			}
			n += tmpBest
		}
		if n <= t && n > best {
			best = n
		}
	}
	return best
}

func max(ls []int) (max int) {
	for _, v := range ls {
		if v > max {
			max = v
		}
	}
	return max
}
// ChooseBestSum : Final solution
var ChooseBestSum func(t, k int, ls []int) int = ChooseBestSumRecur

func main() {
	ls := []int{91, 74, 73, 85, 73, 81, 87}
	fmt.Println(Comb(ls, 3))
	fmt.Println(ChooseBestSum(230, 3, ls))
}
