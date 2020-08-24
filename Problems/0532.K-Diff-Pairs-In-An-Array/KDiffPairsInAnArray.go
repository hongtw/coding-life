package leetcode


// Runtime: 16 ms, faster than 98.98% of Go online submissions
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
		keys := make([]int, 0, len(counter))
		for key := range counter {
			keys = append(keys, key)
		}
		sort.Ints(keys)
		for _, key := range keys {
			if _, isOK := counter[key + k]; isOK {
				res++
			}
		}
	}
	return res
}