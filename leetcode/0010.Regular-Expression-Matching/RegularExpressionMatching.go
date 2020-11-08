package leetcode

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Regular Expression Matching.
func dp(i, j int, s, p string, memo map[[2]int]bool) (ans bool) {
	if ans, isSeen := memo[[2]int{i, j}]; isSeen {
		return ans
	}
	if j == len(p) {
		return i == len(s)
	}
	isFirstOK := i < len(s) && (p[j] == '.' || p[j] == s[i])
	if j+1 < len(p) && p[j+1] == '*' {
		ans = (isFirstOK && dp(i+1, j, s, p, memo)) || dp(i, j+2, s, p, memo)
	} else {
		ans = isFirstOK && dp(i+1, j+1, s, p, memo)
	}
	memo[[2]int{i, j}] = ans
	return ans
}

func isMatch(s string, p string) bool {
	memo := make(map[[2]int]bool)
	return dp(0, 0, s, p, memo)
}
