
from CombinationSum import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([2,3,6,7], 7, [[2,2,3],[7]]),
    ([1, 2], 3, [[1, 1, 1], [1, 2]]),

]

@pytest.mark.parametrize(
    "nums, target, ans", 
    TEST_SUITS
)
def test(nums, target, ans):
    assert SOL.combinationSum(deepcopy(nums), target) == ans
    assert SOL.combinationSumV2(deepcopy(nums), target) == ans