from ContainerWithMostWater import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([1,8,6,2,5,4,8,3,7], 49),
    ([2,3,4,5,18,17,6], 17)
]

@pytest.mark.parametrize(
    "height, ans", 
    deepcopy(TEST_SUITS)
)
def test(height, ans):
    assert SOL.maxArea(height) == ans
