
from ContainsDuplicate import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([1,2,3,1], True),
    ([1,2,3,4], False),
    ([1,1,1,3,3,4,3,2,4,2], True),
]

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, ans):
    assert SOL.containsDuplicate(nums) == ans

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def test_v2(nums, ans):
    assert SOL.containsDuplicate_v2(nums) == ans