
from SortColors import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([2,0,2,1,1,0], sorted([2,0,2,1,1,0])),
    ([2,0,1], sorted([2,0,1])),
    ([0], [0]),
]

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, ans):
    SOL.sortColors(nums)
    assert nums == ans
    