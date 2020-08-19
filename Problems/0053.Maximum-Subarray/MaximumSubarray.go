package leetcode

// Runtime: 4 ms, faster than 97.89% of Go online submissions
func maxSubArray(nums []int) int {
	maxsum := nums[0]
	subsum := 0
	for _, n := range nums {
		subsum += n
		if subsum > maxsum {
			maxsum = subsum
		}
		if subsum < 0 {
			subsum = 0
		}
	}
	return maxsum
}

func maxSubArrayV2(nums []int) int {
	maxsum := nums[0]
	subsum := 0
	for _, n := range nums {
		if subsum+n > n {
			subsum += n
		} else {
			subsum = n
		}
		if subsum > maxsum {
			maxsum = subsum
		}
	}
	return maxsum
}
