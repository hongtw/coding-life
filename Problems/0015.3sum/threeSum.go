package leetcode

import "sort"

// Runtime: 84 ms, faster than 35.65% of Go online submissions
func threeSum(nums []int) [][]int {
	counter := make(map[int]int)
	res := make([][]int, 0)

	for _, n := range nums {
		counter[n]++
	}
	uniqNums := []int{}
	for k := range counter {
		uniqNums = append(uniqNums, k)
	}
	sort.Ints(uniqNums)

	for i := range uniqNums {
		for j := i + 1; j < len(uniqNums); j++ {
			if uniqNums[i]*2+uniqNums[j] == 0 && counter[uniqNums[i]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			} else if uniqNums[i]+uniqNums[j]*2 == 0 && counter[uniqNums[j]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			} else {
				r := 0 - uniqNums[i] - uniqNums[j]
				if counter[r] > 0 && r > uniqNums[j] {
					res = append(res, []int{uniqNums[i], uniqNums[j], r})
				}
			}

		}
	}
	if counter[0] >= 3 {
		res = append(res, []int{0, 0, 0})
	}
	return res
}
