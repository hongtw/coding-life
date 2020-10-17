package leetcode

import "sort"

func combinHelper(candidates []int, target int, prefixRes []int, res *[][]int) {
	for i, n := range candidates {
		if i > 0 && n == candidates[i-1] {
			continue
		}

		if n == target {
			copyPrefix := make([]int, len(prefixRes))
			copy(copyPrefix, prefixRes)
			*res = append(*res, append(copyPrefix, n))
		} else if n < target {
			combinHelper(candidates[i+1:], target-n, append(prefixRes, n), res)
		} else {
			break
		}
	}
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum II.
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	combinHelper(candidates, target, []int{}, &res)
	return res
}
