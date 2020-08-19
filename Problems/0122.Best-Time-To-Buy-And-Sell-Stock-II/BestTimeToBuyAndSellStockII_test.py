
from BestTimeToBuyAndSellStockII import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([7,1,5,3,6,4], 7),
    ([7,6,4,3,1], 0),
    ([1,2,3,4,5], 4),
]

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, ans):
    assert SOL.maxProfit(nums) == ans

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)   
def test_v2(nums, ans):
    assert SOL.maxProfit_v2(nums) == ans