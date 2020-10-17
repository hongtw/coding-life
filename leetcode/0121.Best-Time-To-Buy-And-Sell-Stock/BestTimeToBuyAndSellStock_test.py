
from BestTimeToBuyAndSellStock import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([7,1,5,3,6,4], 5),
    ([7,6,4,3,1], 0),
    ([1,2], 1),
]

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, ans):
    assert SOL.maxProfit(nums) == ans
    