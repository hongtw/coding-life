
from FirstMissingPositive import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([1,2,0], 3),
    ([3,4,-1,1], 2),
    ([7,8,9,11,12], 1),
]

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, ans):
    assert SOL.firstMissingPositive(nums) == ans
    