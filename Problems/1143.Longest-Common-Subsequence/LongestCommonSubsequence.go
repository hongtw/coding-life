package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Subsequence.
func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	for r := 1; r <= len1; r++ {
		for c := 1; c <= len2; c++ {
			if text1[r-1] == text2[c-1] {
				dp[r][c] = dp[r-1][c-1] + 1
			} else {
				dp[r][c] = max(dp[r-1][c], dp[r][c-1])
			}
		}
	}

	return dp[len1][len2]
}

func main() {
	fmt.Println(longestCommonSubsequence("babcde", "ace"))
}
