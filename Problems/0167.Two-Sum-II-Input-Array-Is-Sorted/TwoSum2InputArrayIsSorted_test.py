
from TwoSum2InputArrayIsSorted import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([2,7,11,15], 9, [1, 2]),
]

@pytest.mark.parametrize(
    "nums, target, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, target, ans):
    assert SOL.twoSum(nums, target) == ans