package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Runtime: 32 ms, faster than 66.18% of Go online submissions for Longest Palindromic Subsequence.
func longestPalindromeSubseq(s string) int {
	strLen := len(s)
	dp := make([][]int, strLen)
	for i := 0; i < strLen; i++ {
		dp[i] = make([]int, strLen)
		dp[i][i] = 1
	}

	// i -> 0 and j -> strLen
	for i := strLen; i >= 0; i-- {
		for j := i + 1; j < strLen; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][strLen-1]
}

func isReversedString(s string, strLen int) bool {
	for i, j := 0, strLen-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// Early Return
// Runtime: 8 ms, faster than 97.06% of Go online submissions for Longest Palindromic Subsequence.
func longestPalindromeSubseqV2(s string) int {
	strLen := len(s)
	if isReversedString(s, strLen) {
		return strLen
	}
	dp := make([][]int, strLen)
	for i := 0; i < strLen; i++ {
		dp[i] = make([]int, strLen)
		dp[i][i] = 1
	}

	// i -> 0 and j -> strLen
	for i := strLen; i >= 0; i-- {
		for j := i + 1; j < strLen; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][strLen-1]
}

func main() {
	fmt.Println(longestPalindromeSubseq("aabba"))
}
