package leetcode

import "fmt"

// Runtime: 0 ms, faster than 100.00% of Go online submissions
func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for _, num := range nums {
		for _, sub := range res {
			clone := make([]int, len(sub), len(sub)+1)
			copy(clone, sub)
			res = append(res, append(clone, num))
			// 這邊如果不做 copy ，而是用以下寫法 (直接 append 子slice 並回傳)，
			// 會有預期外的錯誤
			// res = append(res, append(sub, num))
		}
	}
	return res
}