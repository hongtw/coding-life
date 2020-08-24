
from KDiffPairsInAnArray import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([6, 3, 5, 7, 2, 3, 3, 8, 2, 4], 2, 5)
]

@pytest.mark.parametrize(
    "nums, k, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, k, ans):
    assert SOL.findPairs(nums, k) == ans
    assert SOL.findPairsV2(nums, k) == ans
    