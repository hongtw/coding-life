from RemoveElement import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
        ([3,2,2,3], 3, 2,  [2, 2]),
        ([0,1,2,2,3,0,4,2], 2, 5, [0, 1, 3, 0, 4]), 
]

@pytest.mark.parametrize(
    "nums, val, ans, val_nums", 
    deepcopy(TEST_SUITS)
)
def test(nums, val, ans, val_nums):
    assert SOL.removeElement(nums, val) == ans
    for i, n in enumerate(val_nums):
        assert nums[i] == n

@pytest.mark.parametrize(
    "nums, val, ans, val_nums", 
    deepcopy(TEST_SUITS)
)
def test_v2(nums, val, ans, val_nums):
    assert SOL.removeElement_v2(nums, val) == ans
    for i, n in enumerate(val_nums):
        assert nums[i] == n