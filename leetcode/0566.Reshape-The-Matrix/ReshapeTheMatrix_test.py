
from ReshapeTheMatrix import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([[1,2],[3,4]], 1, 4, [[1,2,3,4]]),
    ([[1,2],[3,4]], 2, 4, [[1,2],[3,4]])
]

@pytest.mark.parametrize(
    "nums, r, c, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, r, c, ans):
    assert SOL.matrixReshape(nums, r, c) == ans
    