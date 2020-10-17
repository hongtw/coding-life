from MaximumSubarray import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
        ("Example-1", [-2,1,-3,4,-1,2,1,-5,4], 6),
        ("Example-2", [-1], -1),
]

@pytest.mark.parametrize(
    "example, nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(example, nums, ans):
    assert SOL.maxSubArray(nums) == ans