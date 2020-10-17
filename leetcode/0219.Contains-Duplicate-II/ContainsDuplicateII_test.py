
from ContainsDuplicateII import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([1,2,3,1], 3, True),
    ([1,0,1,1], 1, True),
    ([1,2,3,1,2,3], 2, False),
    ([99, 99], 2, True),
    ([1,2,1], 0, False)
]

@pytest.mark.parametrize(
    "nums, k, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, k, ans):
    assert SOL.containsNearbyDuplicate(nums, k) == ans
    