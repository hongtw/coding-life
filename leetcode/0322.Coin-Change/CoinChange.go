package leetcode

// Runtime: 8 ms, faster than 90.05% of Go online submissions for Coin Change.
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if subRes := 1 + dp[i-coin]; subRes < dp[i] {
				dp[i] = subRes
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// Runtime: 4 ms, faster than 99.73% of Go online submissions for Coin Change
func coinChangeV2(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := range coins {
			coin := coins[j]
			if i-coin < 0 {
				continue
			}
			if subRes := 1 + dp[i-coin]; subRes < dp[i] {
				dp[i] = subRes
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
