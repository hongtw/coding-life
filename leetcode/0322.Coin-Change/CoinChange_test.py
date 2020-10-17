
from CoinChange import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([1, 2, 5], 11, 3),
    ([2], 3, -1),
    ([1], 0, 0),
    ([1], 1, 1),
    ([1], 2, 2),
]

@pytest.mark.parametrize(
    "coins, amount, ans", 
    deepcopy(TEST_SUITS)
)
def test(coins, amount, ans):
    assert SOL.coinChange(coins, amount) == ans
    assert SOL.coinChangeV2(coins, amount) == ans
    