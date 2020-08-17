package leetcode

// Runtime: 0 ms, faster than 100.00% of Go online submissions
func merge(nums1 []int, m int, nums2 []int, n int) {
	pt := m + n - 1
	pt1 := m - 1
	pt2 := n - 1
	for pt2 >= 0 {
		if pt1 < 0 || nums2[pt2] > nums1[pt1] {
			nums1[pt] = nums2[pt2]
			pt2--
		} else {
			nums1[pt] = nums1[pt1]
			pt1--
		}
		pt--
	}
}
