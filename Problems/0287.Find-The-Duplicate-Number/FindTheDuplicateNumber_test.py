
from FindTheDuplicateNumber import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([1,3,4,2,2], 2),
    ([3,1,3,4,2], 3),
    ([2,2,2,2,2], 2),
]

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, ans):
    assert SOL.findDuplicate(nums) == ans

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)  
def testV2(nums, ans):
    assert SOL.findDuplicateV2(nums) == ans