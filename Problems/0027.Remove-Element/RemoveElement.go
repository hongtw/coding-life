package leetcode

// Runtime: 0 ms, faster than 100.00% of Go online submissions
func removeElement(nums []int, val int) int {
	recorder := 0
	for index := range nums {
		if nums[index] != val {
			nums[recorder] = nums[index]
			recorder++
		}
	}
	return recorder
}
