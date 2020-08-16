package leetcode

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

// Runtime: 12 ms, faster than 92.60% of Go online submissions for Container With Most Water.
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxArea := 0
	for l < r {
		if height[l] < height[r] {
			maxArea = max(maxArea, height[l]*(r-l))
			l++
		} else {
			maxArea = max(maxArea, height[r]*(r-l))
			r--
		}
	}
	return maxArea
}
