package leetcode

// Runtime: 12 ms, faster than 96.67% of Go online submissions for Reshape the Matrix.
func matrixReshape(nums [][]int, r int, c int) [][]int {
	h := len(nums)
	w := len(nums[0])
	if h*w != r*c {
		return nums
	}

	res := make([][]int, r, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c, c)
	}

	for i := 0; i < h*w; i++ {
		res[i/c][i%c] = nums[i/w][i%w]
	}
	return res
}
