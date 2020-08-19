package leetcode

// Runtime: 4 ms, faster than 95.13% of Go online submissions
func maxProfit(prices []int) int {
	profits := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profits += prices[i] - prices[i-1]
		}
	}
	return profits
}
