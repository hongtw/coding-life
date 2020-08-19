from typing import List

class Solution:

    # Runtime: 56 ms, faster than 96.03% of Python3 online submissions
    def maxProfit(self, prices: List[int]) -> int:
        length = len(prices)
        buy_price = prices[0]
        profits = 0
        idx = 1
        while idx < length:
            curr_price = prices[idx]
            if buy_price < curr_price:
                if idx + 1 < length:
                    next_price = prices[idx + 1]
                else:
                    next_price = 0
                if curr_price > next_price:
                    profits += curr_price - buy_price
                    buy_price = next_price
                    idx += 1
            else:
                buy_price = curr_price
            idx += 1
        return profits

    def maxProfit_v2(self, prices: List[int]) -> int:
        profits = 0
        for idx in range(1, len(prices)):
            if prices[idx] > prices[idx - 1]:
                profits += prices[idx] - prices[idx - 1]
        return profits
