package leetcode

// Runtime: 12 ms, faster than 100.00% of Go online submissions
func containsDuplicate(nums []int) bool {
	rec := make(map[int]bool, len(nums))
	for _, n := range nums {
		if _, isExist := rec[n]; isExist {
			return true
		}
		rec[n] = true
	}
	return false
}
