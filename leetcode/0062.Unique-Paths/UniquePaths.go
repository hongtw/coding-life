package leetcode

// GolabalCache cache for uniquePaths
var GolabalCache = make(map[[2]int]int)

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths.
func uniquePaths(m int, n int) int {
	tup := [2]int{m, n}
	if res, isOK := GolabalCache[tup]; isOK {
		return res
	} else if m == 1 || n == 1 {
		return 1
	}
	result := uniquePaths(m-1, n) + uniquePaths(m, n-1)
	GolabalCache[tup] = result
	return result
}
