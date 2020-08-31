package leetcode

// Runtime: 8 ms, faster than 60.41% of Go online submissions for Find the Duplicate Number.
func findDuplicate(nums []int) int {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}
	return 1
}

// Runtime: 4 ms, faster than 99.18% of Go online submissions for Find the Duplicate Number.
func findDuplicateV2(nums []int) int {
	pt := 0
	for true {
		if nums[pt] == 0 {
			return pt
		}
		pt, nums[pt] = nums[pt], 0
	}
	return 1
}
