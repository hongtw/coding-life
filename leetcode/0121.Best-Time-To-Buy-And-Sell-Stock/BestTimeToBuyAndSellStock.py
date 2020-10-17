from typing import List

class Solution:

    # Runtime: 56 ms, faster than 97.13% of Python3 online submissions
    def maxProfit(self, prices: List[int]) -> int:
        length = len(prices)
        if length < 2:
            return 0

        max_profit = 0
        buy_price = prices[0]
        for i in range(1, length):
            pr = prices[i]
            if pr < buy_price:
                buy_price = pr
            else:
                profit = pr - buy_price
                if profit > max_profit:
                    max_profit = profit
        return max_profit