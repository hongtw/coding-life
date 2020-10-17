package leetcode

// Runtime: 4 ms, faster than 92.97% of Go online submissions for Move Zeroes.
func moveZeroes(nums []int) {
	length := len(nums)
	pt1 := 0
	pt2 := 1
	for true {
		for pt1 < length && nums[pt1] != 0 {
			pt1++
		}
		for pt2 < length && (pt2 < pt1 || nums[pt2] == 0) {
			pt2++
		}
		if pt1 == length || pt2 == length {
			break
		}
		nums[pt1], nums[pt2] = nums[pt2], nums[pt1]
	}
}
