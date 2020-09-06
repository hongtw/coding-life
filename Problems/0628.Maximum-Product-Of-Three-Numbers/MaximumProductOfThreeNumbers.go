package leetcode

import "fmt"

// Runtime: 28 ms, faster than 100.00% of Go online submissions for Maximum Product of Three Numbers.
func maximumProduct(nums []int) int {
	top1, top2, top3 := -1000, -1000, -1000
	neg1, neg2 := 0, 0

	for _, num := range nums {
		if num > top1 {
			top1, top2, top3 = num, top1, top2
		} else if num > top2 {
			top2, top3 = num, top2
		} else if num > top3 {
			top3 = num
		}

		if num < neg1 {
			neg1, neg2 = num, neg1
		} else if num < neg2 {
			neg2 = num
		}
	}

	fmt.Println("=========", nums, "===", top1, top2, top3, neg1, neg2)
	res1 := top1 * top2 * top3
	res2 := top1 * neg1 * neg2
	if res1 > res2 {
		return res1
	}
	return res2
}
