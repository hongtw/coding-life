package leetcode

// Runtime: 4 ms, faster than 79.63% of Go online submissions for Combination Sum.
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	for i, n := range candidates {
		if n == target {
			res = append(res, []int{n})
		} else if n < target {
			subres := combinationSum(candidates[i:], target-n)
			for _, sub := range subres {
				res = append(res, append([]int{n}, sub...))
			}
		}
	}
	return res
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum.
func combinHelper(candidates []int, target int, prefixRes []int, res *[][]int) {
	for i, n := range candidates {
		if n == target {
			copyPrefix := make([]int, len(prefixRes))
			copy(copyPrefix, prefixRes)
			*res = append(*res, append(copyPrefix, n))
		} else if n < target {
			combinHelper(candidates[i:], target-n, append(prefixRes, n), res)
		}
	}
}

func combinationSumV2(candidates []int, target int) [][]int {
	res := [][]int{}
	combinHelper(candidates, target, []int{}, &res)
	return res
}
