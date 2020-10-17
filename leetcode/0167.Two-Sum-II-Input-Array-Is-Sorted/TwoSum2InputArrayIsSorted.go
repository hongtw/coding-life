package leetcode

// Runtime: 4 ms, faster than 96.00% of Go online submissions
func twoSum(numbers []int, target int) []int {
	rec := make(map[int]int)
	for i, n := range numbers {
		if idx1, isExist := rec[n]; isExist {
			return []int{idx1 + 1, i + 1}
		}
		rec[target-n] = i
	}
	return nil
}
