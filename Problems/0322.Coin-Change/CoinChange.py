from typing import List
from functools import lru_cache


class Solution:

    # Runtime: 1444 ms, faster than 55.25% of Python3 online submissions for Coin Change.
    #  O(kn)
    def coinChange(self, coins: List[int], amount: int) -> int:
        @lru_cache(maxsize=None)
        def dp(n):
            if n == 0: return 0
            if n < 0:  return -1
            res = float('inf')
            for coin in coins:
                sub_cnt = dp(n - coin)
                if sub_cnt == -1:
                    continue
                res = min(res, 1 + sub_cnt)
            return res == float('inf') and -1 or res

        return dp(amount)

    # Runtime: 1340 ms, faster than 67.25% of Python3 online submissions for Coin Change.
    def coinChangeV2(self, coins: List[int], amount: int) -> int:
        dp = [0] + [amount+1] * amount
        for i in range(1, amount+1):
            for coin in coins:
                if i - coin < 0: continue
                dp[i] = min(dp[i], 1 + dp[i - coin])
        return dp[amount] == amount+1 and -1 or dp[amount]


    # Runtime: 820 ms, faster than 95.46% of Python3 online submissions for Coin Change.
    def coinChangeV3(self, coins: List[int], amount: int) -> int:
        dp = [amount+1] * (amount +1)
        dp[0] = 0
        for i in range(1, amount+1):
            for coin in coins:
                if i - coin < 0: continue
                sub_res = 1 + dp[i - coin]
                if sub_res < dp[i]:
                    dp[i] = sub_res
        return dp[amount] == amount+1 and -1 or dp[amount]


if __name__ == "__main__":
    print(Solution().coinChangeV2([1, 2, 5], 11))