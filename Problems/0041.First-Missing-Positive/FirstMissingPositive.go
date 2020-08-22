package leetcode

// Runtime: 0 ms, faster than 100.00% of Go online submissions for First Missing Positive.
func firstMissingPositive(nums []int) int {
	total := len(nums)
	numMap := make(map[int]bool, total)

	for _, num := range nums {
		if num > 0 {
			numMap[num] = true
		}
	}
	for i := 1; i < total+1; i++ {
		if _, isExist := numMap[i]; !isExist {
			return i
		}
	}
	return total + 1
}
