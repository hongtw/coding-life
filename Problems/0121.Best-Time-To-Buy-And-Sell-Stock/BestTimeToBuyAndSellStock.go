package leetcode

// Runtime: 4 ms, faster than 94.81% of Go online submissions
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxprofit := 0
	profit := 0
	buyPrice := prices[0]
	for _, pr := range prices[1:] {
		if pr < buyPrice {
			buyPrice = pr
		} else if profit = pr - buyPrice; profit > maxprofit {
			maxprofit = profit
		}
	}
	return maxprofit
}
