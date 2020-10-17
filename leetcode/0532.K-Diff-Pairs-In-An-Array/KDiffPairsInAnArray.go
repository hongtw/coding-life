package leetcode


// Runtime: 12 ms, faster than 100.00% of Go online submissions for K-diff Pairs in an Array.
func findPairs(nums []int, k int) int {
	res := 0
	if k < 0 {return res}

    counter := make(map[int]int)
	for _, n := range nums {
		counter[n]++
	}
	if  k == 0 {
        for _, v := range counter {
			if v > 1 {
				res += 1
			}
		}
	} else {
		for key := range counter {
			if _, isOK := counter[key + k]; isOK {
				res++
			}
		}
	}
	return res
}