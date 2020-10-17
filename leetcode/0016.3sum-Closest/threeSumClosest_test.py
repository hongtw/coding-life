
from threeSumClosest import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([-1,2,1,-4], 1, 2),
    ([1, 1, 1, 0], -100, 2),
    ([0,2,1,-3], 1, 0),
]

@pytest.mark.parametrize(
    "nums, target, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, target, ans):
    assert SOL.threeSumClosest(nums, target) == ans

@pytest.mark.parametrize(
    "nums, target, ans", 
    deepcopy(TEST_SUITS)
)
def testV2(nums, target, ans):
    assert SOL.threeSumClosestV2(nums, target) == ans