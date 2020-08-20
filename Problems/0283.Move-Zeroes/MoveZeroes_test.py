
from MoveZeroes import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([0,1,0,3,12], [1,3,12,0,0]),
    ([2,1], [2,1]),
    ([4,2,4,0,0,3,0,5,1,0], [4,2,4,3,5,1,0,0,0,0]),
]

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, ans):
    SOL.moveZeroes(nums)
    assert nums == ans
    