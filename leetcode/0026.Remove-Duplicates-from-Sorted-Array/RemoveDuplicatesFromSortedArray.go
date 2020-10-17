package leetcode

// Runtime: 8 ms, faster than 89.71% of Go online submissions
func removeDuplicates(nums []int) int {
	if l := len(nums); l <= 1 {
		return l
	}

	recorder := 0
	for finder := range nums {
		if nums[finder] != nums[recorder] {
			recorder++
			nums[recorder] = nums[finder]
		}
	}
	return recorder + 1
}

// Runtime: 4 ms, faster than 99.45% of Go online submissions
func removeDuplicatesV2(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}

	recorder := 0
	for finder := 0; finder < length; finder++ {
		if nums[finder] != nums[recorder] {
			recorder++
			nums[recorder] = nums[finder]
		}
	}
	return recorder + 1
}
